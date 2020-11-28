package usecase

import (
	"context"
	set "github.com/deckarep/golang-set"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	config "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	"github.com/rocketlaunchr/dataframe-go"
	"gonum.org/v1/gonum/stat"
	"sync"
	"time"
)

type RecommendationSystemUseCase struct {
	RecommendationRepository config.Repository
	UserDataframe            *dataframe.DataFrame
	UserIDRowContainer       map[uint64]int
	MovieNameContainer       []models.Movie
	CorrelationUserMap       map[uint64]map[uint64]float64
	Mu                       *sync.RWMutex
}

type DataframeUserRatingRow map[interface{}]interface{}

func (t DataframeUserRatingRow) ValuesAndWeights(keys []models.Movie) (*[]float64, *[]float64) {
	resultArr := make([]float64, 0)
	weightsArr := make([]float64, 0)
	for _, val := range keys {
		keyValue, ok := t[val.Name]
		if !ok {
			continue
		}
		flVal, ok := keyValue.(float64)
		if !ok {
			weightsArr = append(weightsArr, 0.0)
			resultArr = append(resultArr, flVal)
			continue
		}
		resultArr = append(resultArr, flVal)
		weightsArr = append(weightsArr, 1.0)
	}
	return &resultArr, &weightsArr
}

func (t DataframeUserRatingRow) Values(keys []models.Movie) *[]float64 {
	resultArr := make([]float64, 0)
	for _, val := range keys {
		keyValue, ok := t[val.Name]
		if !ok {
			continue
		}
		flVal, _ := keyValue.(float64)
		resultArr = append(resultArr, flVal)
	}
	return &resultArr
}

func NewRecommendationSystemUseCase(rep config.Repository, sleepTime time.Duration, mutex *sync.RWMutex) *RecommendationSystemUseCase {
	sys := &RecommendationSystemUseCase{
		rep,
		nil,
		make(map[uint64]int),
		make([]models.Movie, 0),
		make(map[uint64]map[uint64]float64),
		mutex,
	}
	readyChan := make(chan bool)
	go sys.UpdateDataframe(sleepTime, readyChan)
	<-readyChan
	return sys
}

func (t *RecommendationSystemUseCase) getUsersMovies(cmpUser, userID uint64, arr *set.Set) {
	t.Mu.RLock()
	cmpUserRow := t.UserIDRowContainer[cmpUser]
	mainUserRow := t.UserIDRowContainer[userID]
	movieNameContainer := t.MovieNameContainer
	t.Mu.RUnlock()
	cmpDfRow := t.UserDataframe.Row(cmpUserRow, false)
	userDfRow := t.UserDataframe.Row(mainUserRow, false)
	for _, val := range movieNameContainer {
		if userDfRow[val.Name] == nil && cmpDfRow[val.Name] != nil && cmpDfRow[val.Name].(float64) > config.RatingBorder {
			(*arr).Add(val.ID)
		}
	}
}

func (t *RecommendationSystemUseCase) UpdateDataframe(sleepTime time.Duration, ch chan bool) {
	firstInit := true
	for {
		dataset, getDatasetErr := t.RecommendationRepository.GetMovieRatingsDataset()
		if getDatasetErr != nil {
			return
		}

		userIDSeries := dataframe.NewSeriesInt64(config.PrimaryUserColumnName, nil)
		userRowCounter := 0
		movieSeriesMap := make(map[uint64]dataframe.Series)
		UserIDRowContainer := make(map[uint64]int)
		MovieNameContainer := make([]models.Movie, 0)

		for _, val := range *dataset {
			if movieSeriesMap[val.MovieID] == nil {
				movieSeriesMap[val.MovieID] = dataframe.NewSeriesFloat64(val.MovieName, nil)
				MovieNameContainer = append(MovieNameContainer, models.Movie{ID: val.MovieID, Name: val.MovieName})
			}

			if userRow, ok := UserIDRowContainer[val.UserID]; ok {
				for movieSeriesMap[val.MovieID].NRows() <= userRow {
					movieSeriesMap[val.MovieID].Append(nil)
				}
				movieSeriesMap[val.MovieID].Update(userRow, float64(val.UserRating))
				continue
			}
			for movieSeriesMap[val.MovieID].NRows() < userRowCounter {
				movieSeriesMap[val.MovieID].Append(nil)
			}
			userIDSeries.Append(val.UserID)
			movieSeriesMap[val.MovieID].Append(float64(val.UserRating))
			UserIDRowContainer[val.UserID] = userRowCounter
			userRowCounter++
		}

		t.UserDataframe = dataframe.NewDataFrame(userIDSeries)
		for _, val := range movieSeriesMap {
			for val.NRows() < userRowCounter {
				val.Append(nil)
			}
			_ = t.UserDataframe.AddSeries(val, nil)
		}
		t.Mu.Lock()
		t.UserIDRowContainer = UserIDRowContainer
		t.MovieNameContainer = MovieNameContainer
		t.Mu.Unlock()

		if firstInit {
			ch <- true
			firstInit = false
		}

		time.Sleep(sleepTime)
	}
}

func (t *RecommendationSystemUseCase) CreateCorrelationArray(userID uint64) error {
	t.Mu.RLock()
	userIDRowContainer := t.UserIDRowContainer
	MovieNameContainer := t.MovieNameContainer
	t.Mu.RUnlock()
	if _, ok := userIDRowContainer[userID]; !ok {
		return models.ErrFooNoRatingInfo
	}

	userIDRatingsRow := t.UserDataframe.Row(userIDRowContainer[userID], false)
	userValuesMap := DataframeUserRatingRow(userIDRatingsRow)
	userValuesArr, weightsArr := userValuesMap.ValuesAndWeights(MovieNameContainer)

	delete(userIDRowContainer, userID)
	resultCorrelationArray := make(map[uint64]float64)

	for key, val := range userIDRowContainer {
		cmpUserRow := t.UserDataframe.Row(val, false)
		cmpUserValuesMap := DataframeUserRatingRow(cmpUserRow)
		cmpUserValuesArr := cmpUserValuesMap.Values(MovieNameContainer)

		corr := stat.Correlation(*userValuesArr, *cmpUserValuesArr, *weightsArr)
		resultCorrelationArray[key] = corr
	}

	t.Mu.Lock()
	t.CorrelationUserMap[userID] = resultCorrelationArray
	t.Mu.Unlock()

	return nil
}

func (t *RecommendationSystemUseCase) MakeMovieRecommendations(userID uint64) (set.Set, error) {
	t.Mu.RLock()
	userCorrelation, ok := t.CorrelationUserMap[userID]
	t.Mu.RUnlock()
	if !ok {
		err := t.CreateCorrelationArray(userID)
		if err != nil {
			return nil, models.ErrFooNoRatingInfo
		}
		t.Mu.RLock()
		userCorrelation, ok = t.CorrelationUserMap[userID]
		t.Mu.RUnlock()
	}
	if !ok {
		return nil, models.ErrFooNoRatingInfo
	}

	correlationUser := dataframe.NewSeriesInt64(config.PrimaryUserColumnName, nil)
	correlationSeries := dataframe.NewSeriesFloat64(config.PrimaryCorrelationColumnName, nil)
	for key, val := range userCorrelation {
		correlationUser.Append(key)
		correlationSeries.Append(val)
	}
	newCorrelationDataframe := dataframe.NewDataFrame(correlationUser, correlationSeries)
	newCorrelationDataframe.Sort(context.Background(), []dataframe.SortKey{
		{
			Key: config.PrimaryCorrelationColumnName, Desc: true,
		},
	})

	FilterFunc := dataframe.FilterDataFrameFn(func(values map[interface{}]interface{}, row, nRows int) (dataframe.FilterAction, error) {
		fl, ok := values[config.PrimaryCorrelationColumnName].(float64)
		if !ok || fl < config.CorrelationBorder {
			return dataframe.DROP, nil
		}
		return dataframe.KEEP, nil
	})
	TmpDF, FilteringErr := dataframe.Filter(context.Background(), newCorrelationDataframe, FilterFunc)
	if FilteringErr != nil {
		return nil, models.ErrFooIncorrectInputInfo
	}
	CastedDataFrame, _ := TmpDF.(*dataframe.DataFrame)

	MovieSet := set.NewSet()
	for i := 0; i < CastedDataFrame.NRows(); i++ {
		dfRow := CastedDataFrame.Row(i, false)
		t.getUsersMovies(uint64(dfRow[config.PrimaryUserColumnName].(int64)), userID, &MovieSet)
	}
	return MovieSet, nil
}

func (t *RecommendationSystemUseCase) GetRecommendedMovieList(userID uint64) (*[]models.Movie, error) {
	MovieIDSet, recommendationErr := t.MakeMovieRecommendations(userID)
	if recommendationErr != nil {
		return t.GetPopularMovies()
	}

	MovieList, GetMovieErr := t.RecommendationRepository.GetRecommendedMovieList(&MovieIDSet)
	if GetMovieErr != nil {
		return nil, models.ErrFooInternalDBErr
	}

	return MovieList, nil
}

func (t *RecommendationSystemUseCase) GetPopularMovies() (*[]models.Movie, error) {
	return t.RecommendationRepository.GetPopularMovies()
}

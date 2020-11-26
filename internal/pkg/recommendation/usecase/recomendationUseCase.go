package usecase

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	dataframe "github.com/rocketlaunchr/dataframe-go"
	"gonum.org/v1/gonum/stat"
	"sync"
	"time"
)

type RecommendationSystemUseCase struct {
	RecommendationRepository recommendation.Repository
	UserDataframe            *dataframe.DataFrame
	UserIDRowContainer       map[uint64]int
	MovieNameContainer		 []string
	CorralationUserMap		 map[uint64]map[uint64]float64
	Mu                       *sync.RWMutex
}

type DataframeUserRatingRow map[interface{}]interface{}

func (t DataframeUserRatingRow) Values(keys []string) (*[]float64, error) {
	resultArr := make([]float64, 0)
	for _, val := range keys{
		keyValue, ok := t[val]
		if !ok {
			continue
		}
		flVal, _ := keyValue.(float64)
		resultArr = append(resultArr, flVal)
	}
	return &resultArr, nil
}

func NewRecommendationSystemUseCase(rep recommendation.Repository, SleepTime time.Duration, mutex *sync.RWMutex) *RecommendationSystemUseCase {
	sys := &RecommendationSystemUseCase{
		rep,
		nil,
		make(map[uint64]int),
		make([]string,0),
		make(map[uint64]map[uint64]float64),
		mutex,
	}
	readyChan := make(chan bool)
	go sys.UpdateDataframe(SleepTime, readyChan)
	<-readyChan
	return sys
}

func (t *RecommendationSystemUseCase) UpdateDataframe(SleepTime time.Duration, ch chan bool) {
	firstInit := true
	for {
		dataset, getDatasetErr := t.RecommendationRepository.GetMovieRatingsDataset()
		if getDatasetErr != nil {
			return
		}

		userIDSeries := dataframe.NewSeriesInt64("UserID", nil)
		userRowCounter := 0
		movieSeriesMap := make(map[uint64]dataframe.Series)
		UserIDRowContainer := make(map[uint64]int)
		MovieNameContainer := make([]string,0)

		for _, val := range *dataset {
			if movieSeriesMap[val.MovieID] == nil {
				movieSeriesMap[val.MovieID] = dataframe.NewSeriesFloat64(val.MovieName, nil)
				MovieNameContainer = append(MovieNameContainer, val.MovieName)
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

		UserDataframe := dataframe.NewDataFrame(userIDSeries)
		for _, val := range movieSeriesMap {
			for val.NRows() < userRowCounter {
				val.Append(nil)
			}
			_ = UserDataframe.AddSeries(val, nil)
		}

		t.Mu.Lock()
		t.UserIDRowContainer = UserIDRowContainer
		t.UserDataframe = UserDataframe
		t.MovieNameContainer = MovieNameContainer
		t.Mu.Unlock()

		if firstInit {
			ch <- true
			firstInit = false
		}

		time.Sleep(SleepTime)
	}
}

func (t *RecommendationSystemUseCase) CreateCorrelationArray(userID uint64) (error) {
	t.Mu.RLock()
	userRatingsDataframe := t.UserDataframe
	userIDRowContainer := t.UserIDRowContainer
	t.Mu.RUnlock()

	userIDRatingsRow := userRatingsDataframe.Row(userIDRowContainer[userID], false)
	userValuesMap := DataframeUserRatingRow(userIDRatingsRow)
	userValuesArr, err := userValuesMap.Values(t.MovieNameContainer)
	if err != nil {
		return models.ErrFooCastErr
	}
	delete(userIDRowContainer, userID)
	resultCorrelationArray := make(map[uint64]float64)

	for key, val := range userIDRowContainer {
		cmpUserRow := userRatingsDataframe.Row(val, false)
		cmpUserValuesMap := DataframeUserRatingRow(cmpUserRow)
		cmpUserValuesArr, err := cmpUserValuesMap.Values(t.MovieNameContainer)
		if err != nil {
			return models.ErrFooCastErr
		}

		corr := stat.Correlation(*userValuesArr, *cmpUserValuesArr, nil)
		resultCorrelationArray[key] = corr
	}

	t.Mu.Lock()
	t.CorralationUserMap[userID] = resultCorrelationArray
	t.Mu.Unlock()

	return nil
}


func (t *RecommendationSystemUseCase) MakeMovieRecommendations(userID uint64) (*[]models.Movie, error){
	t.Mu.RLock()
	userCorrelation, ok := t.CorralationUserMap[userID]
	t.Mu.RUnlock()
	if !ok{
		err := t.CreateCorrelationArray(userID)
		if err != nil{
			return nil, models.ErrFooIncorrectInputInfo
		}
	}

	correlationUser := dataframe.NewSeriesInt64("UserID", nil)
	correlationSeries := dataframe.NewSeriesFloat64("Correlation", nil)
	for key,val := range userCorrelation{
		correlationUser.Append(key)
		correlationSeries.Append(val)
	}
	newCorrelationDataframe := dataframe.NewDataFrame(correlationUser, correlationSeries)
	newCorrelationDataframe.Sort(context.Background(), []dataframe.SortKey{
		{
			Key:"Correlation", Desc: true,
		},
	})

	fmt.Println(newCorrelationDataframe.Table())
	return nil, nil
}










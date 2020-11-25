package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	dataframe "github.com/rocketlaunchr/dataframe-go"
	"sync"

	"gonum.org/v1/gonum/stat"
	"time"
)

type RecommendationSystemUseCase struct{
	RecommendationRepository recommendation.Repository
	UserDataframe *dataframe.DataFrame
	UserIDRowContainer map[uint64]int
	Mu *sync.RWMutex
	IsReady bool
}

type DataframeUserRatingRow map[interface{}]interface{}

func (t DataframeUserRatingRow) Values() (*[]float64,error){
	resultArr := make([]float64,0)
	for _,val := range t{
		if flVal, ok := val.(float64); ok{
			resultArr = append(resultArr, flVal)
			continue
		}
		return nil,models.ErrFooCastErr
	}
	return &resultArr, nil
}


func NewRecommendationSystemUseCase(rep recommendation.Repository, SleepTime time.Duration, mutex *sync.RWMutex) * RecommendationSystemUseCase{
	sys :=  &RecommendationSystemUseCase{
		rep,
		nil,
		make	(map[uint64]int),
		mutex,
		false,
	}
	//sys.SetUpDataframe(SleepTime)
	go sys.SetUpDataframe(SleepTime)
	return sys
}


func (t *RecommendationSystemUseCase) SetUpDataframe(SleepTime time.Duration){
	for {
		dataset, getDatasetErr := t.RecommendationRepository.GetMovieRatingsDataset()
		if getDatasetErr != nil {
			return
		}

		userIDSeries := dataframe.NewSeriesInt64("UserID", nil)
		userRowCounter := 0
		movieSeriesMap := make(map[uint64]dataframe.Series)
		UserIDRowContainer := make(map[uint64]int)

		for _, val := range *dataset {
			if movieSeriesMap[val.MovieID] == nil {
				movieSeriesMap[val.MovieID] = dataframe.NewSeriesFloat64(val.MovieName, nil)
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
		t.IsReady = true
		t.Mu.Unlock()

		time.Sleep(SleepTime)
	}
}


func (t *RecommendationSystemUseCase) MakeRecommendations(userID uint64) (*[]models.Movie, error) {
	// waiting for dataframe initialization
	for !t.IsReady{}

	t.Mu.RLock()
	userRatingsDataframe := t.UserDataframe
	userIDRowContainer := t.UserIDRowContainer
	t.Mu.RUnlock()

	userIDRatingsRow := userRatingsDataframe.Row(userIDRowContainer[userID], false)
	userValuesMap, ok := interface{}(userIDRatingsRow).(DataframeUserRatingRow)
	if !ok{
		return nil,models.ErrFooCastErr
	}

	userValuesArr, err := userValuesMap.Values()
	if err != nil{
		return nil, models.ErrFooCastErr
	}
	delete(userIDRowContainer, userID)


	resultCorrelationArray := make(map[uint64]float64)
	for key, val := range userIDRowContainer{
		cmpUserRow := userRatingsDataframe.Row(val, false)
		cmpUserValuesMap, ok := interface{}(cmpUserRow).(DataframeUserRatingRow)
		if !ok {
			return nil, models.ErrFooCastErr
		}
		cmpUserValuesArr, err := cmpUserValuesMap.Values()
		if err!= nil{
			return nil, models.ErrFooCastErr
		}

		corr := stat.Correlation(*userValuesArr, *cmpUserValuesArr,nil)
		resultCorrelationArray[key] = corr
		fmt.Println(corr)
	}


	fmt.Println(resultCorrelationArray, "kek")
	return nil, nil
}

package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	dataframe "github.com/rocketlaunchr/dataframe-go"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"

)

type RecommendationSystemUseCase struct{
	RecommendationRepository recommendation.Repository
}


func NewRecommendationSystemUseCase(rep recommendation.Repository) * RecommendationSystemUseCase{
	return &RecommendationSystemUseCase{
		rep,
	}
}


func (t *RecommendationSystemUseCase) makeRecommendations() (*[]models.Movie, error) {
	dataset, getDatasetErr := t.RecommendationRepository.GetMovieRatingsDataset()
	if getDatasetErr != nil{
		return nil, models.ErrFooInternalDBErr
	}
	userIDSeries := dataframe.NewSeriesInt64("UserID", nil)
	userIDContainer := make(map[uint64]int)
	userRowCounter := 0
	movieSeriesMap := make(map[uint64]dataframe.Series)

	for _, val := range *dataset{
		if movieSeriesMap[val.MovieID] == nil{
			movieSeriesMap[val.MovieID] = dataframe.NewSeriesInt64(val.MovieName, nil)
		}
		if userRow, ok := userIDContainer[val.UserID]; ok{
			for movieSeriesMap[val.MovieID].NRows() <= userRow{
				movieSeriesMap[val.MovieID].Append(nil)
			}
			movieSeriesMap[val.MovieID].Update(userRow,val.UserRating)
			continue
		}
		for movieSeriesMap[val.MovieID].NRows() < userRowCounter{
			movieSeriesMap[val.MovieID].Append(nil)
		}
		userIDSeries.Append(val.UserID)
		movieSeriesMap[val.MovieID].Append(val.UserRating)
		userIDContainer[val.UserID] = userRowCounter
		userRowCounter ++
	}

	movieDataframe := dataframe.NewDataFrame(userIDSeries)
	for _, val := range movieSeriesMap{
		for val.NRows() < userRowCounter{
			 val.Append(nil)
		}
		_ = movieDataframe.AddSeries(val,nil)
	}
	fmt.Println(movieDataframe.Table())
	ind, _ := movieDataframe.NameToColumn("UserID")
	fmt.Println(movieDataframe.Series[ind])

	stat.CorrelationMatrix()

	return nil, nil
}

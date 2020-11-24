package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	"golang.org/x/tools/go/analysis/passes/stringintconv/testdata/src/a"
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
	dataframeMap := make(map[string]interface{})
	for _,val := range *dataset{
		dataframeMap["UserID"] = val.UserID


	}
	return nil, nil
}

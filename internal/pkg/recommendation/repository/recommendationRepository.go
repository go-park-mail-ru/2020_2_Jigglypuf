package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type RecommendationSystemRepository struct {
	DBConn *sql.DB
}


func (t *RecommendationSystemRepository) GetMovieRatingsDataset() (*[]models.RecommendationDataFrame, error){
	sqlQuery := "SELECT "
	_, resultErr := t.DBConn.Query(sqlQuery)
	if resultErr != nil{
		return nil,models.ErrFooInternalDBErr
	}

	return nil, nil
}
package repository

import (
	"database/sql"
	mapset "github.com/deckarep/golang-set"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/lib/pq"
	"log"
)

type RecommendationSystemRepository struct {
	DBConn *sql.DB
}

func NewRecommendationRepository(connection *sql.DB) *RecommendationSystemRepository{
	return &RecommendationSystemRepository{
		connection,
	}
}


func (t *RecommendationSystemRepository) GetMovieRatingsDataset() (*[]models.RecommendationDataFrame, error) {
	sqlQuery := "SELECT user_id, movie_id, movie_rating,m.moviename, m.rating, m.rating_count FROM rating_history JOIN movie m on m.id = rating_history.movie_id "
	resultSQL, resultErr := t.DBConn.Query(sqlQuery)
	if resultSQL == nil || resultErr != nil || resultSQL.Err() != nil {
		log.Println(resultErr)
		return nil, models.ErrFooInternalDBErr
	}

	resultModelList := make([]models.RecommendationDataFrame, 0)
	resultModel := new(models.RecommendationDataFrame)
	for resultSQL.Next() {
		scanErr := resultSQL.Scan(&resultModel.UserID, &resultModel.MovieID,
			&resultModel.UserRating, &resultModel.MovieName, &resultModel.MovieRating,
			&resultModel.MovieRatingCount)
		if scanErr != nil {
			return nil, models.ErrFooInternalDBErr
		}
		resultModelList = append(resultModelList, *resultModel)
	}

	return &resultModelList, nil
}

func (t *RecommendationSystemRepository) GetRecommendedMovieList(movieIDSet *mapset.Set) (*[]models.Movie, error) {
	SQL := "SELECT v1.ID, v1.MovieName, v1.Description, JSONB_AGG(jsonb_build_object('ID',v3.id,'Name',v3.genre_name)), v1.Duration, v1.Producer, v1.Country,v1.Release_Year, v1.Age_group, v1.Rating, " +
		"v1.Rating_count,JSONB_AGG(jsonb_build_object('ID',v5.ID,'Name', v5.Name, 'Surname', v5.Surname, 'Patronymic', v5.Patronymic, 'Description', v5.Description)), v1.PathToAvatar, v1.pathToSliderAvatar FROM movie v1 " +
		"join movie_genre v2 on v1.id = v2.movie_id " +
		"join genre v3 on (v3.id = v2.genre_id) " +
		"join movie_actors v4 on (v4.movie_id = v1.id) " +
		"join actor v5 on (v5.id = v4.actor_id) " +
		"WHERE v1.ID in $1 GROUP BY v1.ID "
	rows, err := t.DBConn.Query(SQL, pq.Array((*movieIDSet).ToSlice()))
	if err != nil || rows == nil || rows.Err() != nil {
		// TODO log
		log.Println(err)
		return nil, models.ErrFooInternalDBErr
	}
	resultMovieList := make([]models.Movie, 0)
	movieModel := new(models.Movie)
	for rows.Next() {
		ScanErr := rows.Scan(&movieModel.ID, &movieModel.Name, &movieModel.Description,
			&movieModel.GenreList, &movieModel.Duration,
			&movieModel.Producer, &movieModel.Country, &movieModel.ReleaseYear,
			&movieModel.AgeGroup, &movieModel.Rating, &movieModel.RatingCount, &movieModel.ActorList,
			&movieModel.PathToAvatar, &movieModel.PathToSliderAvatar)
		if ScanErr != nil {
			// TODO log
			log.Println(ScanErr)
			return nil, models.ErrFooInternalDBErr
		}
		resultMovieList = append(resultMovieList, *movieModel)
	}

	return &resultMovieList, nil
}

func (t *RecommendationSystemRepository) GetPopularMovies() (*[]models.Movie, error) {
	SQL := "SELECT v1.ID, v1.MovieName, v1.Description, JSONB_AGG(jsonb_build_object('ID',v3.id,'Name',v3.genre_name)), v1.Duration, v1.Producer, v1.Country,v1.Release_Year, v1.Age_group, v1.Rating, " +
		"v1.Rating_count,JSONB_AGG(jsonb_build_object('ID',v5.ID,'Name', v5.Name, 'Surname', v5.Surname, 'Patronymic', v5.Patronymic, 'Description', v5.Description)), v1.PathToAvatar, v1.pathToSliderAvatar FROM movie v1 " +
		"join movie_genre v2 on v1.id = v2.movie_id " +
		"join genre v3 on (v3.id = v2.genre_id) " +
		"join movie_actors v4 on (v4.movie_id = v1.id) " +
		"join actor v5 on (v5.id = v4.actor_id) " +
		"GROUP BY v1.ID ORDER BY v1.Rating_count DESC"

	rows, err := t.DBConn.Query(SQL)
	if err != nil || rows == nil || rows.Err() != nil {
		// TODO log
		log.Println(err)
		return nil, models.ErrFooInternalDBErr
	}
	resultMovieList := make([]models.Movie, 0)
	movieModel := new(models.Movie)
	for rows.Next() {
		ScanErr := rows.Scan(&movieModel.ID, &movieModel.Name, &movieModel.Description,
			&movieModel.GenreList, &movieModel.Duration,
			&movieModel.Producer, &movieModel.Country, &movieModel.ReleaseYear,
			&movieModel.AgeGroup, &movieModel.Rating, &movieModel.RatingCount, &movieModel.ActorList,
			&movieModel.PathToAvatar, &movieModel.PathToSliderAvatar)
		if ScanErr != nil {
			// TODO log
			log.Println(ScanErr)
			return nil, models.ErrFooInternalDBErr
		}
		resultMovieList = append(resultMovieList, *movieModel)
	}

	return &resultMovieList, nil
}

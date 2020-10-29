package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
	"log"
)

type MovieSQLRepository struct {
	DBConnection *sql.DB
}

func NewMovieSQLRepository(connection *sql.DB) *MovieSQLRepository {
	return &MovieSQLRepository{
		connection,
	}
}

func (t *MovieSQLRepository) CreateMovie(movie *models.Movie) error {
	ScanErr := t.DBConnection.QueryRow("INSERT INTO movie (MovieName,Description,Rating,PathToAvatar) VALUES ($1,$2,$3,$4) RETURNING ID",
		movie.Name, movie.Description, movie.Rating, movie.PathToAvatar).Scan(&movie.ID)
	if ScanErr != nil {
		log.Println(ScanErr)
		return ScanErr
	}
	return nil
}

func (t *MovieSQLRepository) UpdateMovie(movie *models.Movie) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("UPDATE movie SET MovieName = $1, Description = $2, Rating = $3, PathToAvatar = $4 WHERE ID = $5",
		movie.Name, movie.Description, movie.Rating, movie.PathToAvatar, movie.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *MovieSQLRepository) GetMovie(id uint64) (*models.Movie, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	resultSQL := t.DBConnection.QueryRow("SELECT ID, MovieName, Description, Rating, Rating_count, PathToAvatar FROM movie WHERE ID = $1", id)
	rowsErr := resultSQL.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}
	resultMovie := new(models.Movie)
	resultErr := resultSQL.Scan(&resultMovie.ID, &resultMovie.Name, &resultMovie.Description, &resultMovie.Rating, &resultMovie.RatingCount, &resultMovie.PathToAvatar)
	if resultErr != nil {
		log.Println(resultErr)
		return nil, resultErr
	}

	return resultMovie, nil
}

func (t *MovieSQLRepository) GetMovieList(limit, page int) (*[]models.Movie, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	resultSQL, DBErr := t.DBConnection.Query("SELECT ID, MovieName, Description, Rating, Rating_count, PathToAvatar FROM movie LIMIT $1 OFFSET $2", limit, page*limit)
	if DBErr != nil {
		log.Println(DBErr)
		return nil, DBErr
	}
	defer resultSQL.Close()

	rowsErr := resultSQL.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}
	movieList := make([]models.Movie, 1)
	for resultSQL.Next() {
		movie := new(models.Movie)
		ScanErr := resultSQL.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Rating, &movie.RatingCount, &movie.PathToAvatar)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, ScanErr
		}
		movieList = append(movieList, *movie)
	}

	return &movieList, nil
}

func (t *MovieSQLRepository) RateMovie(user *models.User, id uint64, rating int64) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("INSERT INTO rating_history (user_id,movie_id,movie_rating) VALUES ($1,$2,$3)",
		user.ID, id, rating)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *MovieSQLRepository) GetRating(user *models.User, id uint64) (int64, error) {
	if t.DBConnection == nil {
		return 0, models.ErrFooNoDBConnection
	}

	resultSQL := t.DBConnection.QueryRow("SELECT movie_rating FROM rating_history WHERE user_id = $1", user.ID)

	rowsErr := resultSQL.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return 0, rowsErr
	}
	var ratingScore int64 = 0
	ScanErr := resultSQL.Scan(&ratingScore)
	if ScanErr != nil {
		log.Println(ScanErr)
		return 0, ScanErr
	}

	return ratingScore, nil
}

func (t *MovieSQLRepository) UpdateMovieRating(movieID uint64, ratingScore int64) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	resultSQL := t.DBConnection.QueryRow("SELECT ID,Rating,Rating_count FROM movie WHERE ID = $1", movieID)
	if resultSQL.Err() != nil {
		return resultSQL.Err()
	}
	var (
		ID          uint64  = 0
		rating      float64 = 0
		RatingCount         = 0
	)
	ScanErr := resultSQL.Scan(&ID, &rating, &RatingCount)
	if ScanErr != nil {
		return ScanErr
	}

	RatingCount++
	var newRating float64 = (rating + float64(ratingScore)) / float64(RatingCount)
	_, RatingDBErr := t.DBConnection.Exec("UPDATE movie SET Rating = $1, Rating_count = $2 WHERE ID = $3",
		newRating, RatingCount, movieID)
	return RatingDBErr
}

func (t *MovieSQLRepository) GetMoviesInCinema(limit, page int) (*[]models.Movie, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	DBRows, DBErr := t.DBConnection.Query("SELECT DISTINCT v1.Movie_id,v2.MovieName,v2.description,v2.rating,v2.rating_count,v2.pathtoavatar FROM movies_in_cinema v1 JOIN movie v2 on(v1.movie_id = v2.id) WHERE rental_start < now() AND rental_end > now()")
	if DBErr != nil {
		log.Println(DBErr)
		return nil, DBErr
	}
	rowsErr := DBRows.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	defer func() {
		if DBRows != nil {
			DBRows.Close()
		}
	}()

	movieList := make([]models.Movie, 1)
	for DBRows.Next() {
		movie := new(models.Movie)
		ScanErr := DBRows.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Rating, &movie.RatingCount, &movie.PathToAvatar)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, ScanErr
		}
		movieList = append(movieList, *movie)
	}

	return &movieList, nil
}

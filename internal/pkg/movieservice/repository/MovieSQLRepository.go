package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
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
	resultSQL := t.DBConnection.QueryRow("SELECT DISTINCT v1.ID, v1.MovieName, v1.Description, JSONB_AGG(jsonb_build_object('ID',v3.id,'Name',v3.genre_name)), v1.Duration, v1.Producer, v1.Country,v1.Release_Year, v1.Age_group, v1.Rating, "+
		"v1.Rating_count,JSONB_AGG(jsonb_build_object('ID',v5.ID,'Name', v5.Name, 'Surname', v5.Surname, 'Patronymic', v5.Patronymic, 'Description', v5.Description)), v1.PathToAvatar, v1.pathToSliderAvatar FROM movie v1 "+
		"join movie_genre v2 on v1.id = v2.movie_id "+
		"join genre v3 on (v3.id = v2.genre_id) "+
		"join movie_actors v4 on (v4.movie_id = v1.id) "+
		"join actor v5 on (v5.id = v4.actor_id) "+
		"WHERE v1.ID = $1 "+
		"GROUP BY v1.ID", id)
	rowsErr := resultSQL.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}
	resultMovie := new(models.Movie)

	resultErr := resultSQL.Scan(&resultMovie.ID, &resultMovie.Name, &resultMovie.Description, &resultMovie.GenreList, &resultMovie.Duration, &resultMovie.Producer, &resultMovie.Country, &resultMovie.ReleaseYear, &resultMovie.AgeGroup, &resultMovie.Rating,
		&resultMovie.RatingCount, &resultMovie.ActorList, &resultMovie.PathToAvatar, &resultMovie.PathToSliderAvatar)
	if resultErr != nil {
		log.Println(resultErr)
		return nil, resultErr
	}

	return resultMovie, nil
}

func (t *MovieSQLRepository) GetMovieList(limit, page int) (*[]models.MovieList, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	resultSQL, DBErr := t.DBConnection.Query("SELECT v1.ID, v1.MovieName, v1.Description, JSONB_AGG(jsonb_build_object('ID',v3.id,'Name',v3.genre_name)), v1.Duration, v1.Producer, v1.Country,v1.Release_Year, v1.Age_group, v1.Rating, "+
		"v1.Rating_count,JSONB_AGG(jsonb_build_object('ID',v5.ID,'Name', v5.Name, 'Surname', v5.Surname, 'Patronymic', v5.Patronymic, 'Description', v5.Description)), v1.PathToAvatar, v1.pathToSliderAvatar FROM movie v1 "+
		"join movie_genre v2 on v1.id = v2.movie_id "+
		"join genre v3 on (v3.id = v2.genre_id) "+
		"join movie_actors v4 on (v4.movie_id = v1.id) "+
		"join actor v5 on (v5.id = v4.actor_id) "+
		"GROUP BY v1.ID "+
		"LIMIT $1 OFFSET $2", limit, page*limit)
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
	movieList := make([]models.MovieList, 0)
	for resultSQL.Next() {
		resultMovie := new(models.MovieList)
		ScanErr := resultSQL.Scan(&resultMovie.ID, &resultMovie.Name, &resultMovie.Description,
			&resultMovie.GenreList, &resultMovie.Duration,
			&resultMovie.Producer, &resultMovie.Country, &resultMovie.ReleaseYear,
			&resultMovie.AgeGroup, &resultMovie.Rating, &resultMovie.RatingCount, &resultMovie.ActorList,
			&resultMovie.PathToAvatar, &resultMovie.PathToSliderAvatar)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, ScanErr
		}
		movieList = append(movieList, *resultMovie)
	}

	return &movieList, nil
}

func (t *MovieSQLRepository) RateMovie(user *models.User, id uint64, rating int64) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("INSERT INTO rating_history (user_id,movie_id,movie_rating) VALUES ($1,$2,$3) on conflict(user_id, movie_id) do update set movie_rating = $3",
		user.ID, id, rating)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *MovieSQLRepository) GetRating(userID uint64, movieID uint64) (int64, error) {
	if t.DBConnection == nil {
		return 0, models.ErrFooNoDBConnection
	}

	resultSQL := t.DBConnection.QueryRow("SELECT movie_rating FROM rating_history WHERE user_id = $1 AND movie_id = $2", userID, movieID)

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
	resultSQL := t.DBConnection.QueryRow("SELECT sum(movie_rating),count(user_id) FROM rating_history WHERE movie_id = $1", movieID)
	if resultSQL.Err() != nil {
		return resultSQL.Err()
	}
	var (
		rating      float64 = 0
		RatingCount         = 0
	)
	ScanErr := resultSQL.Scan(&rating, &RatingCount)
	if ScanErr != nil {
		return ScanErr
	}

	var newRating = (rating) / float64(RatingCount)
	_, RatingDBErr := t.DBConnection.Exec("UPDATE movie SET Rating = $1, Rating_count = $2 WHERE ID = $3",
		newRating, RatingCount, movieID)
	return RatingDBErr
}

func (t *MovieSQLRepository) GetMoviesInCinema(limit, page int, date string, allTime bool) (*[]models.MovieList, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	query := "SELECT DISTINCT v1.ID, v1.MovieName, v1.Description, JSONB_AGG(jsonb_build_object('ID',v3.id,'Name',v3.genre_name)), v1.Duration, v1.Producer, v1.Country,v1.Release_Year, v1.Age_group, v1.Rating, " +
		"v1.Rating_count,JSONB_AGG(jsonb_build_object('ID',v5.ID,'Name', v5.Name, 'Surname', v5.Surname, 'Patronymic', v5.Patronymic, 'Description', v5.Description)), v1.PathToAvatar, v1.pathToSliderAvatar FROM movie v1 " +
		"join movie_genre v2 on v1.id = v2.movie_id " +
		"join genre v3 on (v3.id = v2.genre_id) " +
		"join movie_actors v4 on (v4.movie_id = v1.id) " +
		"join actor v5 on (v5.id = v4.actor_id) "
	if allTime {
		query += "where exists(select 1 from schedule sh where v1.id = sh.movie_id and DATE(sh.premiere_time) > $1) " +
			"group by v1.ID order by v1.Rating_count DESC LIMIT $2 OFFSET $3"
	} else {
		query += "where exists(select 1 from schedule sh where v1.id = sh.movie_id and DATE(sh.premiere_time) = $1) " +
			"group by v1.ID order by v1.Rating_count DESC LIMIT $2 OFFSET $3"
	}

	DBRows, DBErr := t.DBConnection.Query(query, date, limit, page*limit)
	if DBErr != nil || DBRows.Err() != nil {
		log.Println(DBErr)
		return nil, DBErr
	}

	defer func() {
		if DBRows != nil {
			DBRows.Close()
		}
	}()

	movieList := make([]models.MovieList, 0)
	for DBRows.Next() {
		resultMovie := new(models.MovieList)
		ScanErr := DBRows.Scan(&resultMovie.ID, &resultMovie.Name, &resultMovie.Description,
			&resultMovie.GenreList, &resultMovie.Duration,
			&resultMovie.Producer, &resultMovie.Country, &resultMovie.ReleaseYear,
			&resultMovie.AgeGroup, &resultMovie.Rating, &resultMovie.RatingCount, &resultMovie.ActorList,
			&resultMovie.PathToAvatar, &resultMovie.PathToSliderAvatar)
		if ScanErr != nil {
			log.Println(ScanErr)
			return nil, ScanErr
		}
		movieList = append(movieList, *resultMovie)
	}

	return &movieList, nil
}

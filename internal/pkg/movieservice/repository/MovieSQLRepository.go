package repository
import(
	"backend/internal/pkg/models"
	"database/sql"
	"errors"
	"log"
)

type MovieSQLRepository struct{
	DBConnection *sql.DB
}

func NewMovieSQLRepository(connection *sql.DB) *MovieSQLRepository {
	return &MovieSQLRepository{
		connection,
	}
}

func (t *MovieSQLRepository) CreateMovie(movie *models.Movie) error{
	resultSQL, DBErr := t.DBConnection.Exec("INSERT INTO movie (`MovieName`,`Description`,`Rating`,`PathToAvatar`)",
		movie.Name, movie.Description, movie.Rating, movie.PathToAvatar)
	if DBErr != nil{
		log.Println(DBErr)
		return DBErr
	}

	movieID, IDErr := resultSQL.LastInsertId()
	if IDErr != nil{
		return IDErr
	}
	movie.ID = uint64(movieID)
	return nil
}

func (t *MovieSQLRepository) UpdateMovie(movie *models.Movie) error{
	if t.DBConnection == nil{
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("UPDATE movie SET `MovieName` = ?, `Description` = ?, `Rating` = ?, `PathToAvatar` = ? WHERE ID = ?",
		movie.Name, movie.Description, movie.Rating, movie.PathToAvatar, movie.ID)
	if DBErr != nil{
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *MovieSQLRepository) GetMovie(id uint64) (*models.Movie, error){
	if t.DBConnection == nil{
		return nil, errors.New("no database connection")
	}

	resultSQL, DBErr := t.DBConnection.Query("SELECT ID, MovieName, Description, Rating, PathToAvatar FROM movie WHERE ID = ?", id)
	if DBErr != nil{
		log.Println(DBErr)
		return nil,DBErr
	}

	resultMovie := new(models.Movie)
	resultErr := resultSQL.Scan(&resultMovie.ID, &resultMovie.Name, &resultMovie.Description, &resultMovie.Rating, &resultMovie.PathToAvatar)
	if resultErr != nil{
		log.Println(resultErr)
		return nil, resultErr
	}

	return resultMovie, nil
}

func (t *MovieSQLRepository) GetMovieList(limit, page int) (*[]*models.Movie, error){
	if t.DBConnection == nil{
		return nil, errors.New("no database connection")
	}

	resultSQL, DBErr := t.DBConnection.Query("SELECT ID, MovieName, Description, Rating, PathToAvatar FROM movie LIMIT ? OFFSET ?",limit,page*limit)
	if DBErr != nil{
		log.Println(DBErr)
		return nil, DBErr
	}

	movieList := make([]*models.Movie, 1)
	for resultSQL.Next(){
		movie := new(models.Movie)
		ScanErr := resultSQL.Scan(&movie.ID, &movie.Name,&movie.Description, &movie.Rating, &movie.PathToAvatar)
		if ScanErr != nil{
			log.Println(ScanErr)
			return nil,ScanErr
		}
		movieList = append(movieList, movie)
	}

	return &movieList, nil
}

func (t *MovieSQLRepository) RateMovie(user *models.User, id uint64, rating int64) error{
	if t.DBConnection == nil{
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("INSERT INTO rating (`user_id`,`movie_id`,`movie_rating`) VALUES (?,?,?)",
		user.ID,id,rating)
	if DBErr != nil{
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *MovieSQLRepository) GetRating(user *models.User, id uint64) (int64, error){
	if t.DBConnection == nil{
		return 0,errors.New("no database connection")
	}

	resultSQL,DBErr := t.DBConnection.Query("SELECT `movie_rating` FROM rating WHERE `user_id` = ?", user.ID)
	if DBErr != nil{
		return 0, DBErr
	}

	var ratingScore int64 = 0
	ScanErr := resultSQL.Scan(&ratingScore)
	if ScanErr != nil{
		log.Println(ScanErr)
		return 0, ScanErr
	}

	return ratingScore, nil
}

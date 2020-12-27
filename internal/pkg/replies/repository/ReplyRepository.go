package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type ReplyRepository struct {
	dbConnection *sql.DB
}

func NewReplyRepository(dbConnection *sql.DB) *ReplyRepository {
	return &ReplyRepository{
		dbConnection: dbConnection,
	}
}

func (t *ReplyRepository) CreateReply(input *models.ReplyInput, user *models.Profile) error {
	query := "INSERT INTO movie_reply (MovieID, UserID, replyText) VALUES ($1, $2, $3)"
	_, dbErr := t.dbConnection.Exec(query, input.MovieID, user.UserCredentials.ID, input.Text)
	if dbErr != nil {
		return models.ErrFooUniqueFail
	}
	return nil
}

func (t *ReplyRepository) GetMovieReplies(movieID, limit, offset int) ([]models.ReplyModel, error) {
	query := "SELECT v1.ID, v1.MovieID, v1.UserID, v1.replyText, v2.movie_rating FROM movie_reply v1 " +
		"LEFT JOIN rating_history v2 on(v1.UserID = v2.user_id and v1.movieid = v2.movie_id) " +
		"WHERE v1.MovieID = $1 LIMIT $2 OFFSET $3"
	rows, dbErr := t.dbConnection.Query(query, movieID, limit, offset)
	if dbErr != nil {
		return nil, models.ErrFooInternalDBErr
	}

	resultArr := make([]models.ReplyModel, 0)
	for rows.Next() {
		result := new(models.ReplyModel)
		scanErr := rows.Scan(&result.ID, &result.MovieID, &result.User.UserID, &result.Text, &result.UserRating)
		if scanErr != nil {
			return nil, models.ErrFooInternalDBErr
		}
		resultArr = append(resultArr, *result)
	}
	return resultArr, nil
}

func (t *ReplyRepository) UpdateReply(input *models.ReplyUpdateInput, userID uint64) error{
	query := "UPDATE movie_reply SET replyText = $1 WHERE ID = $2 AND userid = $3"
	_, sqlErr := t.dbConnection.Exec(query, input.NewText, input.ReplyID, userID)
	if sqlErr != nil{
		return models.ErrFooIncorrectInputInfo
	}
	return nil
}

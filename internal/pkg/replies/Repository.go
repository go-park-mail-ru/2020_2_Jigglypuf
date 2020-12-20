package replies

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type Repository interface{
	CreateReply(input *models.ReplyInput, user *models.Profile) error
	GetMovieReplies(movieID, limit, offset int) (*[]models.ReplyModel, error)
}

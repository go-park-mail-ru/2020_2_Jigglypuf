package replies

import "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"

type UseCase interface {
	CreateReply(input *models.ReplyInput, UserID uint64) error
	GetMovieReplies(movieID, limit, offset int) (*[]models.ReplyModel, error)
	UpdateReply(input *models.ReplyUpdateInput) error
}

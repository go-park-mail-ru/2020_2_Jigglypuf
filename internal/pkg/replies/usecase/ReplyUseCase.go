package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profile "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
)

type ReplyUseCase struct {
	repository replies.Repository
	profile profile.ProfileServiceClient
	sanitizer  *bluemonday.Policy
}

func NewReplyUseCase(repository replies.Repository, profileClient profile.ProfileServiceClient) *ReplyUseCase {
	return &ReplyUseCase{repository: repository, profile: profileClient, sanitizer: bluemonday.UGCPolicy()}
}


func (t *ReplyUseCase) CreateReply(input *models.ReplyInput, UserID uint64) error{
	utils.SanitizeInput(t.sanitizer, &input.Text)
	if input == nil || input.Text == ""{
		return models.ErrFooIncorrectInputInfo
	}

	prof, err := t.profile.GetProfileByID(context.Background(), &profile.GetProfileByUserIDRequest{UserID: UserID})
	if err != nil{
		return models.ErrFooNoAuthorization
	}

	casted_profile := models.Profile{
		Name: prof.Name,
		Surname: prof.Surname,
		AvatarPath: prof.AvatarPath,
		UserCredentials:&models.User{ID: prof.UserCredentials.UserID},
	}

	return t.repository.CreateReply(input, &casted_profile)
}

func (t *ReplyUseCase) GetMovieReplies(movieID, limit, offset int) (*[]models.ReplyModel, error){
	offset --
	if offset < 0 || limit <= 0{
		return nil, models.IncorrectGetParameters{}
	}

	return t.repository.GetMovieReplies(movieID, limit, offset)
}


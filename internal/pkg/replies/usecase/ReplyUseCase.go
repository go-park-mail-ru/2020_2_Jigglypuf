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
	profile    profile.ProfileServiceClient
	sanitizer  *bluemonday.Policy
}

func NewReplyUseCase(repository replies.Repository, profileClient profile.ProfileServiceClient) *ReplyUseCase {
	return &ReplyUseCase{repository: repository, profile: profileClient, sanitizer: bluemonday.UGCPolicy()}
}

func (t *ReplyUseCase) CreateReply(input *models.ReplyInput, userID uint64) error {
	if input == nil || input.Text == "" {
		return models.ErrFooIncorrectInputInfo
	}
	utils.SanitizeInput(t.sanitizer, &input.Text)

	prof, err := t.profile.GetProfileByID(context.Background(), &profile.GetProfileByUserIDRequest{UserID: userID})
	if err != nil || prof == nil {
		return models.ErrFooNoAuthorization
	}

	castedProfile := models.Profile{
		Name:            prof.Name,
		Surname:         prof.Surname,
		AvatarPath:      prof.AvatarPath,
		UserCredentials: &models.User{ID: userID},
	}

	return t.repository.CreateReply(input, &castedProfile)
}

func (t *ReplyUseCase) GetMovieReplies(movieID, limit, offset int) (*[]models.ReplyModel, error) {
	offset--
	if offset < 0 || limit <= 0 {
		return nil, models.IncorrectGetParameters{}
	}

	return t.repository.GetMovieReplies(movieID, limit, offset)
}


func (t *ReplyUseCase) UpdateReply(input *models.ReplyUpdateInput, userID uint64) error{
	if input.NewText == ""{
		return models.ErrFooIncorrectInputInfo
	}
	return t.repository.UpdateReply(input, userID)
}

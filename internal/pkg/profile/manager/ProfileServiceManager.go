package manager

import (
	"context"
	"database/sql"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	ProfileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/usecase"
)

type ProfileServiceManager struct {
	useCase    *usecase.ProfileUseCase
	authClient authService.AuthenticationServiceClient
}

func NewProfileServiceManager(connect *sql.DB, authClient authService.AuthenticationServiceClient) *ProfileServiceManager {
	return &ProfileServiceManager{
		useCase:    usecase.NewProfileUseCase(repository.NewProfileSQLRepository(connect)),
		authClient: authClient,
	}
}

func (t *ProfileServiceManager) CreateProfile(ctx context.Context, in *ProfileService.CreateProfileRequest) (*ProfileService.Nil, error) {
	err := t.useCase.CreateProfile(&models.Profile{
		Name:            in.Profile.Name,
		Surname:         in.Profile.Surname,
		AvatarPath:      in.Profile.AvatarPath,
		UserCredentials: &models.User{ID: in.Profile.UserCredentials.UserID},
	})
	// TODO monitoring
	return &ProfileService.Nil{}, err
}

func (t *ProfileServiceManager) GetProfile(ctx context.Context, in *ProfileService.GetProfileRequest) (*ProfileService.Profile, error) {
	profile, err := t.useCase.GetProfile(&in.Login)
	if err != nil {
		// TODO monitoring
		return nil, err
	}
	return &ProfileService.Profile{
		Name:    profile.Name,
		Surname: profile.Surname,
		UserCredentials: &ProfileService.UserProfile{
			UserID: profile.UserCredentials.ID,
		},
		AvatarPath: profile.AvatarPath,
	}, nil
}

func (t *ProfileServiceManager) GetProfileByID(ctx context.Context, in *ProfileService.GetProfileByUserIDRequest) (*ProfileService.Profile, error) {
	profile, err := t.useCase.GetProfileViaID(in.UserID)
	if err != nil {
		// TODO monitoring
		return nil, err
	}
	user, userErr := t.authClient.GetUserByID(ctx, &authService.GetUserByIDRequest{UserID: profile.UserCredentials.ID})
	if userErr != nil {
		return nil, userErr
	}
	return &ProfileService.Profile{
		Name:    profile.Name,
		Surname: profile.Surname,
		UserCredentials: &ProfileService.UserProfile{
			Login: user.User.Login,
		},
		AvatarPath: profile.AvatarPath,
	}, nil
}

func (t *ProfileServiceManager) UpdateProfile(ctx context.Context, in *ProfileService.UpdateProfileRequest) (*ProfileService.Nil, error) {
	err := t.useCase.UpdateProfile(in.Profile.UserCredentials.UserID, in.Profile.Name, in.Profile.Surname, in.Profile.AvatarPath)
	// TODO monitoring
	return &ProfileService.Nil{}, err
}

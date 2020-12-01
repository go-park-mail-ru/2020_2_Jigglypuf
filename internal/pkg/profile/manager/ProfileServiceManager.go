package manager

import (
	"context"
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	ProfileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/usecase"
)

type ProfileServiceManager struct{
	useCase *usecase.ProfileUseCase
}

func NewProfileServiceManager(connect *sql.DB) *ProfileServiceManager{
	return &ProfileServiceManager{
		useCase: usecase.NewProfileUseCase(repository.NewProfileSQLRepository(connect)),
	}
}

func (t *ProfileServiceManager) CreateProfile(ctx context.Context, in *ProfileService.CreateProfileRequest) (*ProfileService.Nil, error){
	err := t.useCase.CreateProfile(&models.Profile{
		Name: in.Profile.Name,
		Surname: in.Profile.Surname,
		AvatarPath: in.Profile.AvatarPath,
		UserCredentials: &models.User{ID: in.Profile.UserID},
	})
	// TODO logger
	return &ProfileService.Nil{},err
}

func (t *ProfileServiceManager) GetProfile (ctx context.Context, in *ProfileService.GetProfileRequest) (*ProfileService.Profile, error){
	profile, err := t.useCase.GetProfile(&in.Login)
	if err != nil{
		// TODO logger
		return nil, err
	}
	return &ProfileService.Profile{
		Name: profile.Name,
		Surname: profile.Surname,
		UserID: profile.UserCredentials.ID,
		AvatarPath: profile.AvatarPath,
	}, nil
}

func (t *ProfileServiceManager) GetProfileByID(ctx context.Context, in *ProfileService.GetProfileByUserIDRequest) (*ProfileService.Profile, error){
	profile, err := t.useCase.GetProfileViaID(in.UserID)
	if err != nil{
		// TODO logger
		return nil, err
	}
	return &ProfileService.Profile{
		Name: profile.Name,
		Surname: profile.Surname,
		UserID: profile.UserCredentials.ID,
		AvatarPath: profile.AvatarPath,
	}, nil
}

func (t *ProfileServiceManager) UpdateProfile(ctx context.Context, in *ProfileService.UpdateProfileRequest) (*ProfileService.Nil, error){
	err := t.useCase.UpdateProfile(in.Profile.UserID, in.Profile.Name, in.Profile.Surname, in.Profile.AvatarPath)
	// TODO logger
	return &ProfileService.Nil{}, err
}
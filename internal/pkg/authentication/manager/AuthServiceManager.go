package manager

import (
	"context"
	"database/sql"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
)

type AuthServiceManager struct {
	useCase *usecase.UserUseCase
}

func NewAuthServiceManager(connection *sql.DB, profileService profileService.ProfileServiceClient, salt string) *AuthServiceManager {
	return &AuthServiceManager{
		useCase: usecase.NewUserUseCase(repository.NewAuthSQLRepository(connection),
			profileService, salt),
	}
}

func (t *AuthServiceManager) SignIn(ctx context.Context, in *authService.SignInRequest) (*authService.Response, error){
	userID, err := t.useCase.SignIn(&models.AuthInput{
		Login: in.Data.Login,
		Password: in.Data.Password,
	})
	if err != nil{
		return nil, err
	}

	return &authService.Response{
		UserID: userID,
	}, nil
}

func (t *AuthServiceManager) SignUp(ctx context.Context, in *authService.SignUpRequest) (*authService.Response, error){
	userID, err := t.useCase.SignUp(&models.RegistrationInput{
		Login: in.Data.Login,
		Password: in.Data.Password,
		Name: in.Data.Name,
		Surname: in.Data.Surname,
	})

	if err != nil{
		return nil, err
	}

	return &authService.Response{
		UserID: userID,
	}, nil
}

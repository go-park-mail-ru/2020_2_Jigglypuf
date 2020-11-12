package authserver

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/configs"
	authDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces"
	authRepository "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/repository"
	authUseCase "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	"database/sql"
	"github.com/julienschmidt/httprouter"
)

type AuthService struct {
	AuthenticationDelivery   *authDelivery.UserHandler
	AuthenticationUseCase    *authUseCase.UserUseCase
	AuthenticationRepository interfaces.AuthRepository
	AuthRouter               *httprouter.Router
}

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router {
	authAPIHandler := httprouter.New()

	authAPIHandler.POST(configs.URLPattern+"register/", authHandler.RegisterHandler)
	authAPIHandler.POST(configs.URLPattern+"login/", authHandler.AuthHandler)
	authAPIHandler.POST(configs.URLPattern+"logout/", authHandler.SignOutHandler)

	return authAPIHandler
}

func Start(cookieRepository cookie.Repository, profileRepository profile.Repository, connection *sql.DB) (*AuthService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	authRep := authRepository.NewAuthSQLRepository(connection)
	authCase := authUseCase.NewUserUseCase(authRep, profileRepository, cookieRepository, configs.Salt)
	authHandler := authDelivery.NewUserHandler(authCase)

	router := configureAuthRouter(authHandler)

	return &AuthService{
		authHandler,
		authCase,
		authRep,
		router,
	}, nil
}

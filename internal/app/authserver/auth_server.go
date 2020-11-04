package authserver

import (
	authConfig "backend/internal/pkg/authentication"
	authDelivery "backend/internal/pkg/authentication/delivery"
	authRepository "backend/internal/pkg/authentication/repository"
	authUseCase "backend/internal/pkg/authentication/usecase"
	"backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/profile"
	"database/sql"
	"github.com/julienschmidt/httprouter"
)

type AuthService struct {
	AuthenticationDelivery   *authDelivery.UserHandler
	AuthenticationUseCase    *authUseCase.UserUseCase
	AuthenticationRepository authConfig.AuthRepository
	AuthRouter               *httprouter.Router
}

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router {
	authAPIHandler := httprouter.New()

	authAPIHandler.POST(authConfig.URLPattern+"register/", authHandler.RegisterHandler)
	authAPIHandler.POST(authConfig.URLPattern+"login/", authHandler.AuthHandler)
	authAPIHandler.POST(authConfig.URLPattern+"logout/", authHandler.SignOutHandler)

	return authAPIHandler
}

func Start(cookieRepository cookie.Repository, profileRepository profile.Repository, connection *sql.DB) (*AuthService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	authRep := authRepository.NewAuthSQLRepository(connection)
	authCase := authUseCase.NewUserUseCase(authRep, profileRepository, cookieRepository, authConfig.Salt)
	authHandler := authDelivery.NewUserHandler(authCase)

	router := configureAuthRouter(authHandler)

	return &AuthService{
		authHandler,
		authCase,
		authRep,
		router,
	}, nil
}

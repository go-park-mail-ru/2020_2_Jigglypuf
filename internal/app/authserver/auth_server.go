package authserver

import (
	authConfig "backend/internal/pkg/authentication"
	authDelivery "backend/internal/pkg/authentication/delivery"
	authRepository "backend/internal/pkg/authentication/repository"
	authUseCase "backend/internal/pkg/authentication/usecase"
	"backend/internal/pkg/middleware/cookie"
	"database/sql"
	"errors"
	"github.com/julienschmidt/httprouter"
	"sync"
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

func StartMock(mutex *sync.RWMutex, cookieRepository cookie.Repository) *AuthService {
	authrep := authRepository.NewUserRepository(mutex)
	authCase := authUseCase.NewUserUseCase(authrep, cookieRepository, authConfig.Salt)
	authHandler := authDelivery.NewUserHandler(authCase)

	router := configureAuthRouter(authHandler)

	return &AuthService{
		authHandler,
		authCase,
		authrep,
		router,
	}
}

func Start(cookieRepository cookie.Repository, connection *sql.DB) (*AuthService, error) {
	if connection == nil {
		return nil, errors.New("no database connection")
	}
	authRep := authRepository.NewAuthSQLRepository(connection)
	authCase := authUseCase.NewUserUseCase(authRep, cookieRepository, authConfig.Salt)
	authHandler := authDelivery.NewUserHandler(authCase)

	router := configureAuthRouter(authHandler)

	return &AuthService{
		authHandler,
		authCase,
		authRep,
		router,
	}, nil
}

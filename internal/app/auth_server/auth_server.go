package auth_server


import (
	authConfig "backend/internal/pkg/authentication"
	authDelivery "backend/internal/pkg/authentication/delivery"
	authRepository "backend/internal/pkg/authentication/repository"
	authUseCase "backend/internal/pkg/authentication/usecase"
	"backend/internal/pkg/middleware/cookie"
	"github.com/julienschmidt/httprouter"
	"sync"
)

type AuthService struct{
	AuthenticationDelivery *authDelivery.UserHandler
	AuthenticationUseCase *authUseCase.UserUseCase
	AuthenticationRepository *authRepository.AuthRepository
	AuthRouter *httprouter.Router
}

func configureAuthRouter(authHandler *authDelivery.UserHandler) *httprouter.Router{

	authApiHandler := httprouter.New()

	authApiHandler.POST(authConfig.UrlPattern + "register/",authHandler.RegisterHandler)
	authApiHandler.POST(authConfig.UrlPattern + "login/", authHandler.AuthHandler)
	authApiHandler.POST(authConfig.UrlPattern + "logout/", authHandler.SignOutHandler)

	return authApiHandler
}

func Start(mutex *sync.RWMutex, cookieRepository cookie.Repository) *AuthService{
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
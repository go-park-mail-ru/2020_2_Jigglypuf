package profile_server

import (
	"backend/internal/pkg/authentication"
	profileConfig "backend/internal/pkg/profile"
	profileDelivery "backend/internal/pkg/profile/delivery"
	profileRepository "backend/internal/pkg/profile/repository"
	profileUseCase "backend/internal/pkg/profile/usecase"
	"github.com/julienschmidt/httprouter"
	"sync"
)

type ProfileService struct{
	ProfileRepository *profileRepository.ProfileRepository
	ProfileUseCase *profileUseCase.ProfileUseCase
	ProfileDelivery *profileDelivery.ProfileHandler
	ProfileRouter *httprouter.Router
}

func configureProfileRouter(handler *profileDelivery.ProfileHandler) * httprouter.Router{
	router := httprouter.New()

	router.GET(profileConfig.UrlPattern, handler.GetProfile)
	router.PUT(profileConfig.UrlPattern, handler.UpdateProfile)

	return router
}

func Start(mutex *sync.RWMutex, authRep authentication.AuthRepository) *ProfileService{
	profileRep := profileRepository.NewProfileRepository(mutex, authRep)
	profileUC := profileUseCase.NewProfileUseCase(profileRep)
	profileHandler := profileDelivery.NewProfileHandler(profileUC)

	profileRouter := configureProfileRouter(profileHandler)

	return &ProfileService{
		profileRep,
		profileUC,
		profileHandler,
		profileRouter,
	}
}

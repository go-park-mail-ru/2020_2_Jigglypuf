package profileserver

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/models"
	profileConfig "backend/internal/pkg/profile"
	profileDelivery "backend/internal/pkg/profile/delivery"
	profileRepository "backend/internal/pkg/profile/repository"
	profileUseCase "backend/internal/pkg/profile/usecase"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"sync"
)

type ProfileService struct {
	ProfileRepository profileConfig.Repository
	ProfileUseCase    *profileUseCase.ProfileUseCase
	ProfileDelivery   *profileDelivery.ProfileHandler
	ProfileRouter     *httprouter.Router
}

func configureProfileRouter(handler *profileDelivery.ProfileHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET(profileConfig.URLPattern, handler.GetProfile)
	router.PUT(profileConfig.URLPattern, handler.UpdateProfile)

	return router
}

func StartMock(mutex *sync.RWMutex, authRep authentication.AuthRepository) *ProfileService {
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

func Start(connection *sql.DB) (*ProfileService, error) {
	if connection == nil {
		return nil, models.NoDataBaseConnection
	}
	profileRep := profileRepository.NewProfileSQLRepository(connection)
	profileUC := profileUseCase.NewProfileUseCase(profileRep)
	profileHandler := profileDelivery.NewProfileHandler(profileUC)

	profileRouter := configureProfileRouter(profileHandler)

	return &ProfileService{
		profileRep,
		profileUC,
		profileHandler,
		profileRouter,
	}, nil
}

package profileserver

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileConfig "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	profileDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/delivery"
	profileRepository "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/repository"
	profileUseCase "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/usecase"
	"database/sql"
	"github.com/julienschmidt/httprouter"
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

func Start(connection *sql.DB) (*ProfileService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
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

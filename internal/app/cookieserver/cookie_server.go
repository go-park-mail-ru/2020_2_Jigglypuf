package cookieserver

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	cookieDelivery "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie/delivery"
	cookieRepository "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie/repository"
	cookieUseCase "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie/usecase"
	"github.com/tarantool/go-tarantool"
)

type CookieService struct {
	CookieDelivery   *cookieDelivery.CookieHandler
	CookieUseCase    cookie.UseCase
	CookieRepository cookie.Repository
	DBConnection     *tarantool.Connection
}

// var CookieManager *CookieService

func Start(connection *tarantool.Connection) (*CookieService, error) {
	cookieRep, DBErr := cookieRepository.NewCookieTarantoolRepository(connection)
	if DBErr != nil {
		return nil, DBErr
	}
	cookieUC := cookieUseCase.NewCookieUseCase(cookieRep)
	// cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieUC)
	CookieManager := &CookieService{
		cookieHandler,
		cookieUC,
		cookieRep,
		connection,
	}

	return CookieManager, nil
}

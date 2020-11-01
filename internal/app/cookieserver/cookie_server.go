package cookieserver

import (
	"backend/internal/pkg/middleware/cookie"
	cookieDelivery "backend/internal/pkg/middleware/cookie/delivery"
	cookieRepository "backend/internal/pkg/middleware/cookie/repository"
	"github.com/tarantool/go-tarantool"
)

type CookieService struct {
	CookieDelivery   *cookieDelivery.CookieHandler
	CookieRepository cookie.Repository
	DBConnection     *tarantool.Connection
}

//var CookieManager *CookieService

func Start(connection *tarantool.Connection) (*CookieService, error) {
	cookieRep, DBErr := cookieRepository.NewCookieTarantoolRepository(connection)
	if DBErr != nil {
		return nil, DBErr
	}
	// cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieRep)
	CookieManager := &CookieService{
		cookieHandler,
		cookieRep,
		connection,
	}

	return CookieManager, nil
}

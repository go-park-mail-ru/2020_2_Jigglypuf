package cookieserver

import (
	"backend/internal/pkg/middleware/cookie"
	cookieDelivery "backend/internal/pkg/middleware/cookie/delivery"
	cookieRepository "backend/internal/pkg/middleware/cookie/repository"
	"github.com/tarantool/go-tarantool"
	"sync"
)

type CookieService struct {
	CookieDelivery   *cookieDelivery.CookieHandler
	CookieRepository cookie.Repository
	DBConnection     *tarantool.Connection
}

var CookieManager *CookieService

func Start() (*CookieService, error) {
	connection, DBConnectionErr := tarantool.Connect(cookie.Host+cookie.Port, tarantool.Opts{
		User: cookie.User,
		Pass: cookie.Password,
	})

	if DBConnectionErr != nil {
		return &CookieService{}, DBConnectionErr
	}

	cookieRep, DBErr := cookieRepository.NewCookieTarantoolRepository(connection)
	if DBErr != nil {
		return nil, DBErr
	}
	// cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieRep)
	CookieManager = &CookieService{
		cookieHandler,
		cookieRep,
		connection,
	}
	return CookieManager, nil
}

func StartMock(mutex *sync.RWMutex) *CookieService {
	cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieRep)
	CookieManager = &CookieService{
		cookieHandler,
		cookieRep,
		nil,
	}
	return CookieManager
}

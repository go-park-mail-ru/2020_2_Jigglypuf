package cookieserver

import (
	"backend/internal/pkg/middleware/cookie"
	cookieDelivery "backend/internal/pkg/middleware/cookie/delivery"
	cookieRepository "backend/internal/pkg/middleware/cookie/repository"
	"sync"
)

type CookieService struct {
	CookieDelivery   *cookieDelivery.CookieHandler
	CookieRepository cookie.Repository
}

var CookieManager *CookieService

func Start() (*CookieService, error) {
	cookieRep, DBErr := cookieRepository.NewCookieTarantoolRepository()
	if DBErr != nil {
		return nil, DBErr
	}
	// cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieRep)
	CookieManager = &CookieService{
		cookieHandler,
		cookieRep,
	}
	return CookieManager, nil
}

func StartMock(mutex *sync.RWMutex) *CookieService {
	cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieRep)
	CookieManager = &CookieService{
		cookieHandler,
		cookieRep,
	}
	return CookieManager
}

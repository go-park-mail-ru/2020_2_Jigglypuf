package cookie_server

import(
	cookieDelivery "backend/internal/pkg/middleware/cookie/delivery"
	cookieRepository "backend/internal/pkg/middleware/cookie/repository"
	"sync"
)

type CookieService struct{
	CookieDelivery *cookieDelivery.CookieHandler
	CookieRepository *cookieRepository.CookieRepository
}

var CookieManager *CookieService

func Start(mutex *sync.RWMutex) *CookieService{
	cookieRep := cookieRepository.NewCookieRepository(mutex)
	cookieHandler := cookieDelivery.NewCookieHandler(cookieRep)
	CookieManager = &CookieService{
		cookieHandler,
		cookieRep,
	}
	return CookieManager
}

package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"net/http"
	"time"
)

type CookieUseCase struct {
	Repository cookie.Repository
}

func NewCookieUseCase(repository cookie.Repository) *CookieUseCase {
	return &CookieUseCase{
		repository,
	}
}

func (t *CookieUseCase) CheckCookie(cookieValue *http.Cookie) (uint64, bool) {
	value, cookieErr := t.Repository.GetCookie(cookieValue)
	if cookieErr == nil {
		if time.Now().After(value.Cookie.Expires) {
			return 0, false
		}
		return value.UserID, true
	}
	return 0, false
}

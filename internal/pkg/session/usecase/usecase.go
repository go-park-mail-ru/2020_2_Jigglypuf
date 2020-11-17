package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"net/http"
	"time"
)

type CookieUseCase struct {
	Repository session.Repository
}

func NewCookieUseCase(repository session.Repository) *CookieUseCase {
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

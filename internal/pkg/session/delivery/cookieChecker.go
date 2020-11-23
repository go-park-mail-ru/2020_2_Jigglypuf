package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"net/http"
)

type CookieHandler struct {
	useCase session.UseCase
}

func NewCookieHandler(uc session.UseCase) *CookieHandler {
	return &CookieHandler{
		useCase: uc,
	}
}

func (t *CookieHandler) CheckCookie(r *http.Request) (uint64, bool) {
	cookieValue, err := r.Cookie("session_id")
	if err != nil {
		return 0, false
	}

	return t.useCase.CheckCookie(cookieValue)
}

func (t *CookieHandler) RemoveCookie(cookie *http.Cookie) bool {
	handledErr := t.useCase.RemoveCookie(cookie)
	return handledErr == nil
}

func (t *CookieHandler) SetCookie(cookie *http.Cookie, userID uint64) bool {
	handledErr := t.useCase.SetCookie(cookie, userID)
	return handledErr == nil
}

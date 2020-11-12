package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"net/http"
)

type CookieHandler struct {
	useCase cookie.UseCase
}

func NewCookieHandler(uc cookie.UseCase) *CookieHandler {
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

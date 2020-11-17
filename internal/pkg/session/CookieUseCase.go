//go:generate mockgen -source CookieUseCase.go -destination mock/cookieUC_mock.go -package mock
package session

import "net/http"

type UseCase interface {
	CheckCookie(cookieValue *http.Cookie) (uint64, bool)
}

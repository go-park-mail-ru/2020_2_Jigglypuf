package cookie

import "net/http"

type UseCase interface {
	CheckCookie(cookieValue *http.Cookie) (uint64, bool)
}

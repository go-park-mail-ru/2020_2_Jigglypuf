package session

type isAuth int
type UserID string
type ContextCookieValue string

var (
	SessionCookieName                    = "session_id"
	ContextIsAuthName isAuth             = 401
	ContextUserIDName UserID             = "UserID"
	ContextCookieName ContextCookieValue = "CookieValue"
	DBSpaceName                          = "sessions"
)

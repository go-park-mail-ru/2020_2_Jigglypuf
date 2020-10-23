package cookie

type isAuth int
type UserID string

var (
	SessionCookieName        = "session_id"
	ContextIsAuthName isAuth = 401
	ContextUserIDName UserID = "UserID"
	DBSpaceName              = "sessions"
)

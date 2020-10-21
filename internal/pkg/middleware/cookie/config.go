package cookie

type isAuth int

var (
	SessionCookieName        = "session_id"
	ContextIsAuthName isAuth = 401
)

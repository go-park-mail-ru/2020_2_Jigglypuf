package delivery

import (
	"context"
	"encoding/json"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/promconfig"
	session "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	authService authService.AuthenticationServiceClient
}

func NewUserHandler(authService authService.AuthenticationServiceClient) *UserHandler {
	return &UserHandler{
		authService: authService,
	}
}

func createUserCookie() *http.Cookie {
	return &http.Cookie{
		Name:     session.SessionCookieName,
		Value:    models.RandStringRunes(32),
		Expires:  time.Now().Add(96 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	}
}

func setContextCookie(r *http.Request, userID uint64) context.Context {
	sessionValue := createUserCookie()
	ctx := r.Context()
	ctx = context.WithValue(ctx, session.ContextCookieName, *sessionValue)
	ctx = context.WithValue(ctx, session.ContextUserIDName, userID)
	return ctx
}

// Login godoc
// @Summary login
// @Description login user and get session
// @ID login-user-by-login-data
// @Accept  json
// @Param Login_info body models.AuthInput true "Login information"
// @Success 200
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/auth/login/ [post]
func (t *UserHandler) AuthHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w,promconfig.AuthHandler,status)

	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	authInput := new(models.AuthInput)
	translationError := decoder.Decode(authInput)
	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	userID, err := t.authService.SignIn(r.Context(), &authService.SignInRequest{
		Data: &authService.SignInData{
			Login:    authInput.Login,
			Password: authInput.Password,
		},
	})
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	status = promconfig.StatusSuccess
	*r = *r.WithContext(setContextCookie(r, userID.UserID))
}

// Register godoc
// @Summary Register
// @Description register user and get session
// @ID register-user-by-register-data
// @Accept  json
// @Param Register_info body models.RegistrationInput true "Register information"
// @Success 200
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/auth/register/ [post]
func (t *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w,promconfig.RegisterHandler,status)

	if r.Method != http.MethodPost {
		log.Println("incorrect method")
		models.BadMethodHTTPResponse(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	authInput := new(models.RegistrationInput)

	translationError := decoder.Decode(authInput)
	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	userID, err := t.authService.SignUp(r.Context(), &authService.SignUpRequest{
		Data: &authService.SignUpData{
			Login:    authInput.Login,
			Password: authInput.Password,
			Name:     authInput.Name,
			Surname:  authInput.Surname,
		},
	})
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	status = promconfig.StatusSuccess
	*r = *r.WithContext(setContextCookie(r, userID.UserID))
}

// SignOut godoc
// @Summary SignOut
// @Description SignOut user
// @ID SignOut-user-by-register-data
// @Success 200
// @Failure 405 {object} models.ServerResponse
// @Failure 401 {object} models.ServerResponse
// @Router /api/auth/logout/ [post]
func (t *UserHandler) SignOutHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w,promconfig.SignOutHandler,status)

	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}
	isAuth := r.Context().Value(session.ContextIsAuthName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	cookieValue, cookieErr := r.Cookie(session.SessionCookieName)
	if cookieErr != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	cookieValue.Expires = time.Now().Add(-time.Hour)

	status = promconfig.StatusSuccess
	ctx := r.Context()
	ctx = context.WithValue(ctx, session.ContextCookieName, *cookieValue)
	*r = *r.WithContext(ctx)
}

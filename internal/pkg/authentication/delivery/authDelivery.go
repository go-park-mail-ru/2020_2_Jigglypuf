package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	session "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	useCase interfaces.UserUseCase
}

func NewUserHandler(useCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
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

func setContextCookie(r *http.Request, userID uint64) context.Context{
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

	userID, err := t.useCase.SignIn(authInput)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}
	*r = *r.WithContext(setContextCookie(r, userID))

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

	userID, err := t.useCase.SignUp(authInput)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	*r = *r.WithContext(setContextCookie(r, userID))
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
	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}
	isAuth := r.Context().Value(session.ContextIsAuthName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	cookieValue, _ := r.Cookie(session.SessionCookieName)
	expiredCookie, useCaseError := t.useCase.SignOut(cookieValue)
	if useCaseError != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, session.ContextCookieName, *expiredCookie)
	*r = *r.WithContext(ctx)
}

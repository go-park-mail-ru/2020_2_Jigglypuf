package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type UserHandler struct {
	useCase interfaces.UserUseCase
}

func NewUserHandler(useCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}

// Login godoc
// @Summary login
// @Description login user and get cookie
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

	cookie, err := t.useCase.SignIn(authInput)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}
	http.SetCookie(w, cookie)
}

// Register godoc
// @Summary Register
// @Description register user and get cookie
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

	cookie, err := t.useCase.SignUp(authInput)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	http.SetCookie(w, cookie)
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
	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	cookieValue, _ := r.Cookie(cookieService.SessionCookieName)
	expiredCookie, useCaseError := t.useCase.SignOut(cookieValue)
	if useCaseError != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	http.SetCookie(w, expiredCookie)
}

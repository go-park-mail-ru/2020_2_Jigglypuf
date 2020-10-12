package delivery

import (
	"backend/internal/pkg/authentication"
	cookieService "backend/internal/pkg/cookie"
	"backend/internal/pkg/models"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	useCase authentication.UserUseCase
}

func NewUserHandler(useCase authentication.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}

func (t *UserHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {
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

func (t *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
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

func (t *UserHandler) SignOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}

	cookie, cookieError := r.Cookie(cookieService.SessionCookieName)
	if cookieError != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	expiredCookie, useCaseError := t.useCase.SignOut(cookie)
	if useCaseError != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	http.SetCookie(w, expiredCookie)
}

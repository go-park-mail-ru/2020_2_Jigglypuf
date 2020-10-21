package delivery

import (
	cookieService "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/profile"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type ProfileHandler struct {
	useCase profile.UseCase
}

type SavingError struct{}

func (t SavingError) Error() string {
	return "Cannot save the file!"
}

func NewProfileHandler(useCase profile.UseCase) *ProfileHandler {
	return &ProfileHandler{
		useCase: useCase,
	}
}

func SaveAvatarImage(image multipart.File, handler *multipart.FileHeader, fileErr error) (string, error) {
	returnPath := profile.MediaPath
	if fileErr != nil {
		return "", SavingError{}
	}

	defer image.Close()
	fileName := handler.Filename + time.Now().String()
	f, saveErr := os.OpenFile(profile.SavingPath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if saveErr != nil {
		return "", SavingError{}
	}

	defer f.Close()
	_, copyingError := io.Copy(f, image)

	if copyingError != nil {
		return "", SavingError{}
	}
	returnPath += fileName
	return returnPath, nil
}

func (t *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	cookieValue, _ := r.Cookie(cookieService.SessionCookieName)
	requiredProfile, profileError := t.useCase.GetProfileViaCookie(cookieValue)

	if profileError != nil {
		models.BadBodyHTTPResponse(&w, profileError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseProfile, responseErr := json.Marshal(requiredProfile)
	if responseErr != nil {
		models.BadBodyHTTPResponse(&w, responseErr)
		return
	}

	_, _ = w.Write(responseProfile)
}

func (t *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}

	translationError := r.ParseMultipartForm(32 << 20)

	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}
	cookieValue, _ := r.Cookie(cookieService.SessionCookieName)

	profileUpdate, profileError := t.useCase.GetProfileViaCookie(cookieValue)

	if profileError != nil {
		models.BadBodyHTTPResponse(&w, profileError)
		return
	}

	avatarPath, savingErr := SaveAvatarImage(r.FormFile(profile.AvatarFormName))
	if savingErr != nil {
		avatarPath = ""
	}

	err := t.useCase.UpdateProfile(profileUpdate, r.FormValue(profile.NameFormName), r.FormValue(profile.SurnameFormName), avatarPath)

	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	// http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

package delivery

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"backend/models"
	"net/http"
	"os"
	"backend/profile"
)


type ProfileHandler struct {
	useCase profile.ProfileUseCase
}

type SavingError struct{}

func (t SavingError) Error()string{
	return "Cannot save the file!"
}

func NewProfileHandler (useCase profile.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{
		useCase: useCase,
	}
}

func SaveAvatarImage( file multipart.File, handler *multipart.FileHeader, fileErr error )( string, error ){
	returnPath := "/media/"
	if fileErr != nil{
		return "", SavingError{}
	}

	defer file.Close()
	f, saveErr := os.OpenFile("../../media/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if saveErr != nil {
		return "", SavingError{}
	}

	defer f.Close()
	_, copyingError := io.Copy(f, file)

	if copyingError != nil{
		return "", SavingError{}
	}
	returnPath += handler.Filename
	return returnPath, nil
}

func (t *ProfileHandler) GetProfile( w http.ResponseWriter, r *http.Request ) {
	defer r.Body.Close()

	if r.Method != http.MethodGet{
		models.BadMethodHttpResponse(&w)
		return
	}

	w.Header().Set("Content-Type","application/json")

	cookieValue, cookieErr := r.Cookie("session_id")

	if cookieErr != nil{
		models.UnauthorizedHttpResponse(&w)
		return
	}

	requiredProfile, profileError := t.useCase.GetProfileViaCookie(cookieValue)

	if profileError != nil{
		models.BadBodyHTTPResponse(&w, profileError)
		return
	}


	w.WriteHeader(http.StatusOK)
	responseProfile,responseErr := json.Marshal(requiredProfile)
	if responseErr != nil{
		models.BadBodyHTTPResponse(&w, responseErr)
		return
	}

	_, _ = w.Write([]byte(responseProfile))
}

func (t *ProfileHandler) UpdateProfile( w http.ResponseWriter, r *http.Request ) {
	defer r.Body.Close()

	w.Header().Set("Content-Type","application/json")

	if r.Method != http.MethodPost{
		models.BadMethodHttpResponse(&w)
		return
	}

	translationError := r.ParseMultipartForm(32 << 20)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	cookieValue, cookieErr := r.Cookie("session_id")

	if cookieErr != nil{
		models.UnauthorizedHttpResponse(&w)
		return
	}

	profileUpdate, profileError := t.useCase.GetProfileViaCookie(cookieValue)

	if profileError != nil{
		models.BadBodyHTTPResponse(&w, profileError)
		return
	}

	avatarPath, savingErr := SaveAvatarImage(r.FormFile("avatar"))
	if savingErr != nil{
		avatarPath = ""
	}

	err := t.useCase.UpdateProfile(profileUpdate,r.FormValue("name"),r.FormValue("surname"), avatarPath)

	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	//http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
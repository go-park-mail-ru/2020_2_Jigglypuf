package delivery

import (
	"encoding/json"
	"models"
	"net/http"
	"profile"
)


type ProfileHandler struct {
	useCase profile.ProfileUseCase
}


func NewProfileHandler (useCase profile.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{
		useCase: useCase,
	}
}


func (t *ProfileHandler) CreateProfile( w http.ResponseWriter, r *http.Request ) {
	defer r.Body.Close()

	w.Header().Set("Content-Type","application/json")

	decoder := json.NewDecoder(r.Body)
	profileToCreate := new(models.Profile)
	translationError := decoder.Decode(profileToCreate)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	//cookie, err := t.useCase.CreateProfile(profileToCreate)
	err := t.useCase.CreateProfile(profileToCreate)

	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	//http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func (t *ProfileHandler) GetProfile( w http.ResponseWriter, r *http.Request ) {
	defer r.Body.Close()

	w.Header().Set("Content-Type","application/json")

	decoder := json.NewDecoder(r.Body)
	profileGet := new(models.Profile)
	translationError := decoder.Decode(profileGet)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	//cookie, err := t.useCase.GetProfile(&profileGet.Login.Username)
	_, err := t.useCase.GetProfile(&profileGet.Login.Username)

	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	//http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func (t *ProfileHandler) UpdateProfile( w http.ResponseWriter, r *http.Request ) {
	defer r.Body.Close()

	w.Header().Set("Content-Type","application/json")

	decoder := json.NewDecoder(r.Body)
	profileUpdate := new(models.Profile)
	translationError := decoder.Decode(profileUpdate)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	//cookie, err := t.useCase.UpdateProfile(profileUpdate)
	err := t.useCase.UpdateProfile(profileUpdate)

	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	//http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
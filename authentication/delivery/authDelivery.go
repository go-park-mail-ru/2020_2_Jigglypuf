package delivery

import(
	"authentication/usecase"
	"models"
	"net/http"
	"encoding/json"
)


type UserHandler struct {
	useCase usecase.UserUseCase
}


func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler{
	return &UserHandler{
		useCase: *useCase,
	}
}


func BadBodyHTTPResponse(w *http.ResponseWriter, err error){
	response, err := json.Marshal(models.ServerResponse{
		StatusCode: 401,
		Response:  []byte(err.Error()),
	})

	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write(response)
}



func (t *UserHandler) AuthHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	authInput := new(models.AuthInput)
	translationError := decoder.Decode(authInput)

	if translationError != nil{
		BadBodyHTTPResponse(&w, translationError)
		return
	}

	cookie, err := t.useCase.SignIn(authInput)

	if err != nil{
		BadBodyHTTPResponse(&w, err)
		return
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}


func (t *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	w.Header().Set("Content-Type","application/json")

	decoder := json.NewDecoder(r.Body)
	authInput := new(models.RegistrationInput)
	translationError := decoder.Decode(authInput)

	if translationError != nil{
		BadBodyHTTPResponse(&w, translationError)
		return
	}

	cookie, err := t.useCase.SignUp(authInput)

	if err != nil{
		BadBodyHTTPResponse(&w, err)
		return
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
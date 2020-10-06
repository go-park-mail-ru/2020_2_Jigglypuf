package delivery

import (
	"cinemaService"
	"encoding/json"
	"models"
	"net/http"
)

type CinemaHandler struct{
	cinemaUseCase cinemaService.CinemaUseCase
}

func NewCinemaHandler(useCase cinemaService.CinemaUseCase) *CinemaHandler{
	return &CinemaHandler{
		cinemaUseCase: useCase,
	}
}

func (t *CinemaHandler) CreateCinema(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost{
		models.BadMethodHttpResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	cinema := new(models.Cinema)
	translationError := decoder.Decode(cinema)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	createCinemaError := t.cinemaUseCase.CreateCinema(cinema)

	if createCinemaError != nil{
		models.BadBodyHTTPResponse(&w, createCinemaError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *CinemaHandler) GetCinema(w http.ResponseWriter, r* http.Request){
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")



	if r.Method != http.MethodGet{
		models.BadMethodHttpResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	cinema := new(models.SearchCinema)
	translationError := decoder.Decode(cinema)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	result, GetCinemaError := t.cinemaUseCase.GetCinema(&cinema.Name)

	if GetCinemaError != nil{
		models.BadBodyHTTPResponse(&w, GetCinemaError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
	}

	_, _ = w.Write([]byte(response))
}


func (t *CinemaHandler) GetCinemaList(w http.ResponseWriter, r* http.Request){
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")



	if r.Method != http.MethodGet{
		models.BadMethodHttpResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	cinema := new(models.GetCinemaList)

	translationError := decoder.Decode(cinema)
	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}
	result, GetCinemaError := t.cinemaUseCase.GetCinemaList(cinema.Limit, cinema.Page)

	if GetCinemaError != nil{
		models.BadBodyHTTPResponse(&w, GetCinemaError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
	}

	_, _ = w.Write([]byte(response))
}
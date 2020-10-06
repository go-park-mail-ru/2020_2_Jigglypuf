package delivery

import (
	"encoding/json"
	"models"
	"movieService"
	"net/http"
)

type MovieHandler struct{
	movieUseCase movieService.MovieUseCase
}

func NewMovieHandler(usecase movieService.MovieUseCase) *MovieHandler{
	return &MovieHandler{
		movieUseCase: usecase,
	}
}

func (t *MovieHandler) GetMovieList(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	if r.Method != http.MethodGet{
		models.BadMethodHttpResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	movie := new( models.GetMovieList )
	translationError := decoder.Decode(movie)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	resultArray, err := t.movieUseCase.GetMovieList(movie)

	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(resultArray)
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
	}

	_, _ = w.Write([]byte(response))
}

func (t *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	if r.Method != http.MethodGet{
		models.BadMethodHttpResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	movie := new( models.SearchMovie )
	translationError := decoder.Decode(movie)

	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	result, err := t.movieUseCase.GetMovie(movie.Name)

	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
	}

	_, _ = w.Write([]byte(response))
}
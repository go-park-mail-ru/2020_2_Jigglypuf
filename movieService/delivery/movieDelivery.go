package delivery

import (
	"encoding/json"
	"models"
	"movieService"
	"net/http"
	"strconv"
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
	Limit := r.URL.Query()["limit"]
	Page := r.URL.Query()["page"]
	if len(Limit) == 0 || len(Page) == 0{
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}
	limit,limitErr := strconv.Atoi(Limit[0])
	page, pageErr := strconv.Atoi(Page[0])

	if limitErr != nil || pageErr != nil{
		models.BadBodyHTTPResponse(&w, limitErr)
		return
	}

	resultArray, err := t.movieUseCase.GetMovieList(limit,page)

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

	name := r.URL.Query()["name"]

	if len(name) == 0{
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	result, err := t.movieUseCase.GetMovie(name[0])

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
package delivery

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/movieService"
	"backend/internal/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type MovieHandler struct{
	movieUseCase   movieService.MovieUseCase
	userRepository authentication.AuthRepository
}

func NewMovieHandler(usecase movieService.MovieUseCase, userRepository authentication.AuthRepository) *MovieHandler {
	return &MovieHandler{
		movieUseCase: usecase,
		userRepository: userRepository,
	}
}

func (t *MovieHandler) GetMovieList(w http.ResponseWriter, r *http.Request){
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
		return
	}

	_, _ = w.Write([]byte(response))
}

func (t *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request){
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
		return
	}

	_, _ = w.Write([]byte(response))
}

func (t *MovieHandler) RateMovie(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	if r.Method != http.MethodPost{
		models.BadMethodHttpResponse(&w)
		return
	}

	w.Header().Set("Content-Type","application/json")

	cookieValue, cookieErr := r.Cookie("session_id")

	if cookieErr != nil{
		models.UnauthorizedHttpResponse(&w)
		return
	}

	reqUser, userError := t.userRepository.GetUserViaCookie(cookieValue)
	if userError != nil{
		models.UnauthorizedHttpResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	movie := new(models.RateMovie)
	translationError := decoder.Decode(movie)
	if translationError != nil{
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	RateErr := t.movieUseCase.RateMovie( reqUser,movie.Name,movie.Rating )
	if RateErr != nil{
		models.BadBodyHTTPResponse(&w, RateErr)
		return
	}
}

func (t *MovieHandler) GetMovieRating(w http.ResponseWriter, r *http.Request){
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

	reqUser, userError := t.userRepository.GetUserViaCookie(cookieValue)
	if userError != nil{
		models.UnauthorizedHttpResponse(&w)
		return
	}

	name := r.URL.Query()["name"]

	if len(name) == 0{
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	result, RatingErr := t.movieUseCase.GetRating(reqUser, name[0])
	if RatingErr != nil{
		models.BadBodyHTTPResponse(&w, RatingErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	_, _ = w.Write([]byte(response))
}
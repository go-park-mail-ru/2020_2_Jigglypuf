package delivery

import (
	"backend/internal/pkg/authentication"
	cookieService "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/movieservice"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	movieUseCase   movieservice.MovieUseCase
	userRepository authentication.AuthRepository
}

func NewMovieHandler(usecase movieservice.MovieUseCase, userRepository authentication.AuthRepository) *MovieHandler {
	return &MovieHandler{
		movieUseCase:   usecase,
		userRepository: userRepository,
	}
}

func (t *MovieHandler) GetMovieList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	Limit := r.URL.Query()["limit"]
	Page := r.URL.Query()["page"]
	if len(Limit) == 0 || len(Page) == 0 {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}
	limit, limitErr := strconv.Atoi(Limit[0])
	page, pageErr := strconv.Atoi(Page[0])

	if limitErr != nil || pageErr != nil {
		models.BadBodyHTTPResponse(&w, limitErr)
		return
	}

	resultArray, err := t.movieUseCase.GetMovieList(limit, page)

	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(resultArray)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	_, _ = w.Write(response)
}

func (t *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	name := params.ByName(movieservice.GetMovieID)
	integerName, castErr := strconv.Atoi(name)
	if castErr != nil || len(name) == 0 {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	result, err := t.movieUseCase.GetMovie(uint64(integerName))

	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	_, _ = w.Write(response)
}

func (t *MovieHandler) RateMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	UserID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || UserID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	reqUser, userError := t.userRepository.GetUserByID(UserID.(uint64))
	if userError != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	movie := new(models.RateMovie)
	translationError := decoder.Decode(movie)
	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	RateErr := t.movieUseCase.RateMovie(reqUser, movie.ID, movie.Rating)
	if RateErr != nil {
		models.BadBodyHTTPResponse(&w, RateErr)
		return
	}
}

func (t *MovieHandler) GetMovieRating(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	UserID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || UserID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	reqUser, userError := t.userRepository.GetUserByID(UserID.(uint64))
	if userError != nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	name := r.URL.Query()["id"]
	integerName, castErr := strconv.Atoi(name[0])
	if castErr != nil || len(name) == 0 {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	result, RatingErr := t.movieUseCase.GetRating(reqUser, uint64(integerName))
	if RatingErr != nil {
		models.BadBodyHTTPResponse(&w, RatingErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	_, _ = w.Write(response)
}

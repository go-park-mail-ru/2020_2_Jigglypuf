package delivery

import (
	cookieService "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/movieservice"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	movieUseCase movieservice.MovieUseCase
}

func getQueryLimitPageArgs(r *http.Request) (int, int, error) {
	Limit := r.URL.Query()[movieservice.LimitQuery]
	Page := r.URL.Query()[movieservice.PageQuery]
	if len(Limit) == 0 || len(Page) == 0 {
		return 0, 0, models.IncorrectGetParameters{}
	}
	limit, limitErr := strconv.Atoi(Limit[0])
	page, pageErr := strconv.Atoi(Page[0])

	if limitErr != nil || pageErr != nil {
		return 0, 0, limitErr
	}
	return limit, page, nil
}

func NewMovieHandler(usecase movieservice.MovieUseCase) *MovieHandler {
	return &MovieHandler{
		movieUseCase: usecase,
	}
}

// Movie godoc
// @Summary GetMovieList
// @Description Get movie list
// @ID movie-list-id
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {array} models.MovieList
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /movie/ [get]
func (t *MovieHandler) GetMovieList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	limit, page, queryErr := getQueryLimitPageArgs(r)
	if queryErr != nil {
		models.BadBodyHTTPResponse(&w, queryErr)
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

// Movie godoc
// @Summary GetMovie
// @Description Get movie
// @ID movie-id
// @Param id path int true "movie id"
// @Success 200 {object} models.Movie
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /movie/{id}/ [get]
func (t *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	name := vars[movieservice.GetMovieID]
	integerName, castErr := strconv.Atoi(name)
	if castErr != nil || len(name) == 0 {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}
	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	UserID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) {
		isAuth = false
		UserID = uint64(0)
	}

	result, err := t.movieUseCase.GetMovie(uint64(integerName), isAuth.(bool), UserID.(uint64))

	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)

	_, _ = w.Write(response)
}

// Movie godoc
// @Summary RateMovie
// @Description Rate movie
// @ID movie-rate-id
// @Accept  json
// @Param Login_info body models.RateMovie true "Login information"
// @Success 200
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No authorization"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Router /movie/rate/ [post]
func (t *MovieHandler) RateMovie(w http.ResponseWriter, r *http.Request) {
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

	decoder := json.NewDecoder(r.Body)
	movie := new(models.RateMovie)
	translationError := decoder.Decode(movie)
	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	RateErr := t.movieUseCase.RateMovie(UserID.(uint64), movie.ID, movie.Rating)
	if RateErr != nil {
		models.BadBodyHTTPResponse(&w, RateErr)
		return
	}
}

// Movie godoc
// @Summary Get movies in cinema
// @Description Returns movie that in the cinema
// @ID movie-in-cinema-id
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {array} models.MovieList
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No authorization"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Router /movie/actual/ [get]
func (t *MovieHandler) GetMoviesInCinema(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	limit, page, queryErr := getQueryLimitPageArgs(r)
	if queryErr != nil {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	movieList, movieErr := t.movieUseCase.GetMoviesInCinema(limit, page)
	if movieErr != nil {
		models.BadBodyHTTPResponse(&w, movieErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(movieList)

	_, _ = w.Write(response)
}

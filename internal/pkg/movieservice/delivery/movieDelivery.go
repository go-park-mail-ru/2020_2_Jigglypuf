package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/promconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/gorilla/mux"
	"io/ioutil"
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
		return 0, 0, models.IncorrectGetParameters{}
	}
	return limit, page, nil
}

func NewMovieHandler(useCase movieservice.MovieUseCase) *MovieHandler {
	return &MovieHandler{
		movieUseCase: useCase,
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
// @Router /api/movie/ [get]
func (t *MovieHandler) GetMovieList(w http.ResponseWriter, r *http.Request) {
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w, promconfig.GetMovieList, &status)

	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	limit, page, queryErr := getQueryLimitPageArgs(r)
	if queryErr != nil {
		models.BadBodyHTTPResponse(&w, queryErr)
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

	status = promconfig.StatusSuccess
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
// @Router /api/movie/{id}/ [get]
func (t *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w, promconfig.GetMovie, &status)

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
	response, _ := result.MarshalJSON()

	status = promconfig.StatusSuccess
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
// @Router /api/movie/rate/ [post]
func (t *MovieHandler) RateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w, promconfig.RateMovie, &status)

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

	inputBuf, inputErr := ioutil.ReadAll(r.Body)
	if inputErr != nil {
		models.BadBodyHTTPResponse(&w, models.ErrFooIncorrectInputInfo)
	}
	movie := new(models.RateMovie)
	translationError := movie.UnmarshalJSON(inputBuf)
	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	RateErr := t.movieUseCase.RateMovie(UserID.(uint64), movie.ID, movie.Rating)
	if RateErr != nil {
		models.BadBodyHTTPResponse(&w, RateErr)
		return
	}

	status = promconfig.StatusSuccess
}

// Movie godoc
// @Summary Get movies in cinema
// @Description Returns movie that in the cinema
// @ID movie-in-cinema-id
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Param date query int false "date in format 2006-01-02"
// @Success 200 {array} models.MovieList
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No authorization"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Router /api/movie/actual/ [get]
func (t *MovieHandler) GetActualMovies(w http.ResponseWriter, r *http.Request) {
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w, promconfig.GetActualMovies, &status)

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
	date := r.URL.Query()[schedule.DateQueryParamName]

	movieList, movieErr := t.movieUseCase.GetActualMovies(limit, page, date)
	if movieErr != nil {
		models.BadBodyHTTPResponse(&w, movieErr)
		return
	}

	status = promconfig.StatusSuccess
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(movieList)

	_, _ = w.Write(response)
}

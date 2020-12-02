package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/promconfig"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CinemaHandler struct {
	cinemaUseCase cinemaservice.UseCase
}

func NewCinemaHandler(useCase cinemaservice.UseCase) *CinemaHandler {
	return &CinemaHandler{
		cinemaUseCase: useCase,
	}
}

// Cinema godoc
// @Summary GetCinema
// @Description Get cinema
// @ID cinema-id
// @Param id path int true "cinema id param"
// @Success 200 {object} models.Cinema
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/cinema/{id}/ [get]
func (t *CinemaHandler) GetCinema(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w,promconfig.GetCinema,&status)

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	name := params.ByName(cinemaservice.CinemaIDParam)
	integerName, castErr := strconv.Atoi(name)
	if castErr != nil || len(name) == 0 {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	result, GetCinemaError := t.cinemaUseCase.GetCinema(uint64(integerName))

	if GetCinemaError != nil {
		models.BadBodyHTTPResponse(&w, GetCinemaError)
		return
	}

	status = promconfig.StatusSuccess
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

// Cinema godoc
// @Summary GetCinemaList
// @Description Get cinema list
// @ID cinema-list-id
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {array} models.Cinema
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/cinema/ [get]
func (t *CinemaHandler) GetCinemaList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w,promconfig.GetCinemaList,&status)

	w.Header().Set("Content-Type", "application/json")

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
	result, GetCinemaError := t.cinemaUseCase.GetCinemaList(limit, page)

	if GetCinemaError != nil {
		models.BadBodyHTTPResponse(&w, GetCinemaError)
		return
	}

	status = promconfig.StatusSuccess

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

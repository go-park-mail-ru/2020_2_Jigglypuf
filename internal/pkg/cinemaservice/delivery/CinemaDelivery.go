package delivery

import (
	cinemaConfig "backend/internal/pkg/cinemaservice"
	"backend/internal/pkg/cinemaservice"
	"backend/internal/pkg/models"
	"encoding/json"
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

func (t *CinemaHandler) GetCinema(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	name := params.ByName(cinemaConfig.CinemaIdParam)
	integerName, castErr := strconv.Atoi(name);
	if castErr != nil || len(name) == 0 {
		models.BadBodyHTTPResponse(&w, models.IncorrectGetParameters{})
		return
	}

	result, GetCinemaError := t.cinemaUseCase.GetCinema(uint64(integerName))

	if GetCinemaError != nil {
		models.BadBodyHTTPResponse(&w, GetCinemaError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
	}

	_, _ = w.Write(response)
}

func (t *CinemaHandler) GetCinemaList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

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

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(result)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
	}

	_, _ = w.Write(response)
}

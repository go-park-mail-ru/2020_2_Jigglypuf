package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/promconfig"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HallDelivery struct {
	UseCase hallservice.UseCase
}

func NewHallDelivery(useCase hallservice.UseCase) *HallDelivery {
	return &HallDelivery{
		UseCase: useCase,
	}
}

// Hall godoc
// @Summary Get hall structure
// @Description Get cinema hall placement structure
// @ID hall-id
// @Param id path int true "hall id param"
// @Success 200 {object} models.CinemaHall
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Failure 500 {object} models.ServerResponse
// @Router /api/hall/{id}/ [get]
func (t *HallDelivery) GetHallStructure(w http.ResponseWriter, r *http.Request) {
	status := promconfig.StatusErr
	defer promconfig.SetRequestMonitoringContext(w, promconfig.GetHallStructure, &status)
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	hallID := vars[hallservice.HallIDPathName]
	if hallID == "" {
		models.BadBodyHTTPResponse(&w, models.ErrFooIncorrectInputInfo)
		return
	}

	hallItem, hallErr := t.UseCase.GetHallStructure(hallID)
	if hallErr != nil {
		models.BadBodyHTTPResponse(&w, models.ErrFooIncorrectInputInfo)
		return
	}
	log.Println("hall:", hallItem.PlaceConfig)
	status = promconfig.StatusSuccess
	outputBuf, _ := hallItem.MarshalJSON()

	_, _ = w.Write(outputBuf)
}

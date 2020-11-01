package delivery

import (
	"backend/internal/pkg/hallService"
	"backend/internal/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HallDelivery struct{
	UseCase hallService.UseCase
}

func NewHallDelivery(useCase hallService.UseCase)*HallDelivery{
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
// @Router /hall/{id}/ [get]
func (t *HallDelivery) GetHallStructure(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	hallID := vars[hallService.HallIDPathName]
	if hallID == ""{
		models.BadBodyHTTPResponse(&w,models.ErrFooIncorrectInputInfo)
		return
	}

	hallItem, hallErr := t.UseCase.GetHallStructure(hallID)
	if hallErr != nil{
		models.BadBodyHTTPResponse(&w,models.ErrFooIncorrectInputInfo)
		return
	}

	outputBuf, castErr := json.Marshal(hallItem)
	if castErr != nil{
		models.InteralErrorHttpResponse(&w)
		return
	}
	_,_ = w.Write(outputBuf)
}

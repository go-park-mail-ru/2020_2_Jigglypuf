package delivery

import (
	cookieService "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/ticketService"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type TicketDelivery struct{
	useCase ticketService.UseCase
}

func NewTicketDelivery(useCase ticketService.UseCase) *TicketDelivery{
	return &TicketDelivery{
		useCase: useCase,
	}
}


func (t *TicketDelivery) BuyTicket(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	r.Body.Close()
	ticketItem := new(models.Ticket)
	decodeErr:=decoder.Decode(ticketItem)
	if decodeErr != nil{
		models.BadBodyHTTPResponse(&w,decodeErr)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil {
		if ticketItem.Username == ""{
			models.BadBodyHTTPResponse(&w,models.ErrFooNoLoginInfo)
			return
		}
	}

	buyErr := t.useCase.BuyTicket(ticketItem, userID.(uint64))
	if buyErr != nil{
		models.BadBodyHTTPResponse(&w, buyErr)
		return
	}
}

func (t *TicketDelivery) GetUserTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	ticketList,getTicketErr := t.useCase.GetUserTickets(userID.(uint64))
	if getTicketErr != nil{
		models.BadBodyHTTPResponse(&w, getTicketErr)
		return
	}
	outputBuf, castErr := json.Marshal(ticketList)
	if castErr != nil{
		models.InteralErrorHttpResponse(&w)
		return
	}
	_,_ = w.Write(outputBuf)
}

func (t *TicketDelivery) GetUsersSimpleTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	vars := mux.Vars(r)
	ticketID := vars[ticketService.TicketIDQuery]

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	ticketItem, ticketErr := t.useCase.GetSimpleTicket(userID.(uint64), ticketID)
	if ticketErr != nil{
		models.BadBodyHTTPResponse(&w,ticketErr)
		return
	}

	outputBuf, outputErr := json.Marshal(ticketItem)
	if outputErr != nil{
		models.InteralErrorHttpResponse(&w)
		return
	}

	_,_ = w.Write(outputBuf)
}

func (t *TicketDelivery) GetHallScheduleTickets(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}
	vars := mux.Vars(r)
	scheduleID := vars[ticketService.ScheduleIDName]
	ticketList, ticketErr := t.useCase.GetHallScheduleTickets(scheduleID)
	if ticketErr != nil{
		models.BadBodyHTTPResponse(&w, ticketErr)
		return
	}

	outputBuf, castErr := json.Marshal(ticketList)
	if castErr != nil{
		models.InteralErrorHttpResponse(&w)
		return
	}

	_, _ = w.Write(outputBuf)
}
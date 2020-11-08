package delivery

import (
	cookieService "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/ticketservice"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type TicketDelivery struct {
	useCase ticketservice.UseCase
}

func NewTicketDelivery(useCase ticketservice.UseCase) *TicketDelivery {
	return &TicketDelivery{
		useCase: useCase,
	}
}

// Ticket godoc
// @Summary Buy ticket
// @Description Buys ticket by schedule ID and place to authenticated user or by e-mail
// @ID buy-ticket-id
// @Accept  json
// @Param Ticket_info body models.TicketInput true "Ticket info"
// @Success 200
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Router /ticket/buy/ [post]
func (t *TicketDelivery) BuyTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	ticketItem := new(models.TicketInput)
	decodeErr := decoder.Decode(ticketItem)
	if decodeErr != nil {
		models.BadBodyHTTPResponse(&w, decodeErr)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil {
		if ticketItem.Login == "" {
			models.BadBodyHTTPResponse(&w, models.ErrFooNoLoginInfo)
			return
		}
	}

	buyErr := t.useCase.BuyTicket(ticketItem, userID.(uint64))
	if buyErr != nil {
		models.BadBodyHTTPResponse(&w, buyErr)
		return
	}
}

// Ticket godoc
// @Summary Get user ticket list
// @Description Get user ticket list
// @ID get-ticket-list-id
// @Success 200 {array} models.Ticket
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No auth"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Failure 500 {object} models.ServerResponse "Internal err"
// @Router /ticket/ [get]
func (t *TicketDelivery) GetUserTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	ticketList, getTicketErr := t.useCase.GetUserTickets(userID.(uint64))
	if getTicketErr != nil {
		models.BadBodyHTTPResponse(&w, getTicketErr)
		return
	}
	outputBuf, castErr := json.Marshal(ticketList)
	if castErr != nil {
		models.InternalErrorHTTPResponse(&w)
		return
	}
	_, _ = w.Write(outputBuf)
}

// Ticket godoc
// @Summary Get user ticket
// @Description Get user ticket by id
// @ID get-ticket-id
// @Param id path int true "ticket id"
// @Success 200 {object} models.Ticket
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No auth"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Failure 500 {object} models.ServerResponse "Internal err"
// @Router /ticket/{id}/ [get]
func (t *TicketDelivery) GetUsersSimpleTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ticketID := vars[ticketservice.TicketIDQuery]

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	ticketItem, ticketErr := t.useCase.GetSimpleTicket(userID.(uint64), ticketID)
	if ticketErr != nil {
		models.BadBodyHTTPResponse(&w, ticketErr)
		return
	}

	outputBuf, outputErr := json.Marshal(ticketItem)
	if outputErr != nil {
		models.InternalErrorHTTPResponse(&w)
		return
	}

	_, _ = w.Write(outputBuf)
}

// Ticket godoc
// @Summary Get schedule hall ticket list
// @Description Get schedule hall ticket list by id
// @ID get-schedule-ticket-list-id
// @Param id path int true "schedule_id"
// @Success 200 {array} models.TicketPlace
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Failure 500 {object} models.ServerResponse "Internal err"
// @Router /ticket/schedule/{id}/ [get]
func (t *TicketDelivery) GetHallScheduleTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	scheduleID := vars[ticketservice.ScheduleIDName]
	ticketList, ticketErr := t.useCase.GetHallScheduleTickets(scheduleID)
	if ticketErr != nil {
		models.BadBodyHTTPResponse(&w, ticketErr)
		return
	}

	outputBuf, castErr := json.Marshal(ticketList)
	if castErr != nil {
		models.InternalErrorHTTPResponse(&w)
		return
	}

	_, _ = w.Write(outputBuf)
}

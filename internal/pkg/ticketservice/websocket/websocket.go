package websocket

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type TicketWebSocketDelivery struct {
	upgrade websocket.Upgrader
	pool *Pool
}

func NewTicketWebSocketDelivery() *TicketWebSocketDelivery {
	t := &TicketWebSocketDelivery{
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		pool: NewPool(),
	}
	go t.pool.Start()
	return t
}


func (t *TicketWebSocketDelivery) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := t.upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, ErrFooWebSocketUpgradeError
	}
	return conn, nil
}

func (t *TicketWebSocketDelivery) ServeWS(w http.ResponseWriter, r *http.Request) {
	//isAuth := r.Context().Value(session.ContextIsAuthName)
	//userID := r.Context().Value(session.ContextUserIDName)
	//if isAuth == nil || userID == nil || !isAuth.(bool) {
	//	models.UnauthorizedHTTPResponse(&w)
	//	return
	//}
	input, ok := mux.Vars(r)[ticketservice.ScheduleIDName]
	parsedScheduleID, err := strconv.Atoi(input)
	if !ok || err != nil {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}

	conn, err := t.Upgrade(w, r)
	if err != nil {
		models.InternalErrorHTTPResponse(&w)
		return
	}

	client := &Client{
		Conn: conn,
		ScheduleID: uint64(parsedScheduleID),
		Pool: t.pool,
	}

	t.pool.Register <- client
	client.Read()
}

func (t *TicketWebSocketDelivery) ServeBuys(scheduleID uint64, p []models.TicketPlace){
	for _, place := range p{
		t.pool.BroadCast <- Message{ScheduleID: scheduleID, PlaceConfig: WSTicketPlace{Type: Busy, Place: place.Place,
			Row: place.Row}}
	}
	return
}

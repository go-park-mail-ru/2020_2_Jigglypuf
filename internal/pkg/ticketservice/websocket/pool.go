package websocket

import (
	"fmt"
	"log"
)

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[uint64]map[*Client]bool),
		History:    make(map[uint64][]Message, 0),
		BroadCast:  make(chan Message),
	}
}

func (t *Pool) Start() {
	for {
		select {
		case client := <-t.Register:
			if t.Clients[client.ScheduleID] == nil {
				t.Clients[client.ScheduleID] = make(map[*Client]bool)
			}
			t.Clients[client.ScheduleID][client] = true
			break
		case client := <-t.Unregister:
			delete(t.Clients[client.ScheduleID], client)
			break
		case msg := <-t.BroadCast:
			t.ConfigureHistory(&msg)
			fmt.Println("im before sending")
			for client, _ := range t.Clients[msg.ScheduleID] {
				if err := client.Conn.WriteJSON(t.History[msg.ScheduleID]); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func (t *Pool) ConfigureHistory(msg *Message){
	var val int64 = -1
	for index, _ := range t.History[msg.ScheduleID] {
		if t.History[msg.ScheduleID][index].PlaceConfig.Place == msg.PlaceConfig.Place &&
			t.History[msg.ScheduleID][index].PlaceConfig.Row == msg.PlaceConfig.Row {
			val = int64(index)
		}
	}
	if val == -1 && (msg.PlaceConfig.Type == InProcess || msg.PlaceConfig.Type == Busy){
		t.History[msg.ScheduleID] = append(t.History[msg.ScheduleID], *msg)
	}else if val != -1{
		t.History[msg.ScheduleID] = append(t.History[msg.ScheduleID][:val], t.History[msg.ScheduleID][(val+1):]...)
		if msg.PlaceConfig.Type == Busy {
			t.History[msg.ScheduleID] = append(t.History[msg.ScheduleID], *msg)
		}
	}
}
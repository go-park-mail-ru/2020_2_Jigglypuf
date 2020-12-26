package websocket

import (
	"fmt"
	"log"
)

func NewPool() *Pool {
	t := &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[uint64]map[*Client]bool),
		History:    make(map[uint64][]Message, 0),
		BroadCast:  make(chan Message),
	}
	return t
}

func (t *Pool) Start() {
	for {
		select {
		case client := <-t.Register:
			if t.Clients[client.ScheduleID] == nil {
				t.Clients[client.ScheduleID] = make(map[*Client]bool)
			}
			t.Clients[client.ScheduleID][client] = true
			t.Send(client.ScheduleID)
			break
		case client := <-t.Unregister:
			delete(t.Clients[client.ScheduleID], client)
			i := 0
			for i < len(t.History[client.ScheduleID]){
				val := t.History[client.ScheduleID][i]
				if val.PlaceConfig.Type == InProcess && val.Client.Conn == client.Conn{
					fmt.Println("iansdinasooinasoidasd")
					if len(t.History[client.ScheduleID]) > i{
						t.History[client.ScheduleID] = append(t.History[client.ScheduleID][:i],
							t.History[client.ScheduleID][(i+1):]...)
					}else{
						t.History[client.ScheduleID] = t.History[client.ScheduleID][:i]
					}
					i --
				}
				i ++
			}
			fmt.Println("history", t.History[client.ScheduleID])
			t.Send(client.ScheduleID)
			break
		case msg := <-t.BroadCast:
			t.ConfigureHistory(&msg)
			fmt.Println("im before sending")
			t.Send(msg.ScheduleID)
			break
		}
	}
}

func (t *Pool) Send(scheduleID uint64){
	if t.History[scheduleID] == nil{
		t.History[scheduleID] = make([]Message, 0)
	}
	for client, _ := range t.Clients[scheduleID] {
		if err := client.Conn.WriteJSON(t.History[scheduleID]); err != nil {
			log.Println(err)
			return
		}
	}
}

func (t *Pool) ConfigureHistory(msg *Message){
	var val int64 = -1
	for index, _ := range t.History[msg.ScheduleID] {
		if t.History[msg.ScheduleID][index].PlaceConfig.Place == msg.PlaceConfig.Place &&
			t.History[msg.ScheduleID][index].PlaceConfig.Row == msg.PlaceConfig.Row{
			val = int64(index)
		}
	}
	if val == -1 && (msg.PlaceConfig.Type == InProcess || msg.PlaceConfig.Type == Busy){
		t.History[msg.ScheduleID] = append(t.History[msg.ScheduleID], *msg)
	}else if val != -1 && t.History[msg.ScheduleID][val].PlaceConfig.Type != Busy{
		if len(t.History[msg.ScheduleID]) > int(val){
			t.History[msg.ScheduleID] = append(t.History[msg.ScheduleID][:val], t.History[msg.ScheduleID][(val+1):]...)
		}else{
			t.History[msg.ScheduleID] = t.History[msg.ScheduleID][:val]
		}
		if msg.PlaceConfig.Type != Free {
			t.History[msg.ScheduleID] = append(t.History[msg.ScheduleID], *msg)
		}
	}
}
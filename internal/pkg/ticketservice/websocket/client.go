package websocket

import (
	"encoding/json"
	"log"
)

func (t *Client) Read(){
	defer func() {
		t.Pool.Unregister <- t
		_ = t.Conn.Close()
	}()

	for {
		msType, Body, err := t.Conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		input := new(WSTicketPlace)
		err = json.Unmarshal(Body, input)
		if err != nil || input.Type == Busy{
			log.Println(err)
			return
		}
		msg := Message{
			Type: msType,
			ScheduleID: t.ScheduleID,
			PlaceConfig: *input,
			Client: t,
		}
		t.Pool.BroadCast <- msg
		log.Println("Received msg from ws", msg.PlaceConfig)
	}
}

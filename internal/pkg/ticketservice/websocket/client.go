package websocket

import (
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
		log.Println("Received msg from ws", Body, msType)
	}
}

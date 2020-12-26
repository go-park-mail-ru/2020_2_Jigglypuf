package websocket

import (
	"errors"
	"github.com/gorilla/websocket"
)

type Place int

const(
	InProcess Place = iota
	Busy
	Free
)

var(
	ErrFooWebSocketUpgradeError = errors.New("cannot create websocket connection")
)

type Client struct {
	ID string
	ScheduleID uint64
	Conn *websocket.Conn
	Pool *Pool
}


type Message struct {
	Type int `json:"-"`
	ScheduleID uint64
	PlaceConfig WSTicketPlace
	Client *Client `json:"-"`
}

type WSTicketPlace struct {
	Type Place
	Place int
	Row int
}


type Pool struct {
	Register chan *Client
	Unregister chan *Client
	Clients map[uint64]map[*Client]bool
	BroadCast chan Message
	History map[uint64][]Message
}
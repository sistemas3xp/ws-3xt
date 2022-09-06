package ws

import (
	"ws-3xt/src/models"

	socketio "github.com/googollee/go-socket.io"
)

type WS struct {
	server *socketio.Server
}

func NewWS(server *socketio.Server) *WS {
	return &WS{server}
}

func (ws *WS) Run() {
	var rooms []models.Room
	events := NewEvents(&rooms, ws.server)
	ws.server.OnConnect("/", events.Connection)
	ws.server.OnEvent("/", "create-room", events.CreateRoom)
	ws.server.OnEvent("/", "join-room", events.JoinRoom)
	ws.server.OnEvent("/", "play", events.Play)
	ws.server.OnDisconnect("/", events.Disconnection)
}

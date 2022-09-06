package ws

import (
	"fmt"
	"ws-3xt/src/models"
	"ws-3xt/src/utils"

	socketio "github.com/googollee/go-socket.io"
	"github.com/lithammer/shortuuid/v4"
)

type Events struct {
	rooms  *[]models.Room
	server *socketio.Server
}

func NewEvents(rooms *[]models.Room, server *socketio.Server) *Events {
	return &Events{rooms, server}
}

func (e *Events) Connection(socket socketio.Conn) error {
	fmt.Printf("Client %s connected\n", socket.ID())
	return nil
}

func (e *Events) Disconnection(socket socketio.Conn, disconnect_reason string) {
	if len(socket.Rooms()) == 2 {
		e.server.BroadcastToRoom("/", socket.Rooms()[1], "player-retreated")
		room, _ := utils.GetRoomById(*e.rooms, socket.Rooms()[1])
		if room.User1.ID == socket.ID() {
			room.User1 = room.User2
			room.User2 = models.User{}
		} else {
			room.User2 = models.User{}
		}
		socket.LeaveAll()
	}
	fmt.Printf("Client %s disconnected. %s", socket.ID(), disconnect_reason)

}
func (e *Events) CreateRoom(socket socketio.Conn) {
	roomID := shortuuid.New()
	room := models.Room{
		ID:    roomID,
		User1: models.User{ID: socket.ID()},
		User2: models.User{},
	}

	*e.rooms = append(*e.rooms, room)
	socket.Join(roomID)
	socket.Emit("room-created", room)
}

func (e *Events) JoinRoom(socket socketio.Conn, roomID string) {
	room, matches := utils.GetRoomById(*e.rooms, roomID)
	type roomReturn struct {
		code int
		room models.Room
	}
	if !matches {
		socket.Emit("join-status", roomReturn{
			code: 404,
			room: models.Room{},
		})
		return
	}

	if room.User2 != (models.User{}) {
		socket.Emit("join-status", roomReturn{
			code: 401,
			room: models.Room{},
		})
		return
	}

	room.User2 = models.User{
		ID: socket.ID(),
	}
	socket.Join(room.ID)
	socket.Emit("join-status", roomReturn{
		code: 200,
		room: room,
	})
	e.server.BroadcastToRoom("/", room.ID, "start-game")
}

func (e *Events) Play(socket socketio.Conn, board models.Board) {
	e.server.BroadcastToRoom("/", socket.Rooms()[1], "play", board)
}

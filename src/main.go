package main

import (
	"log"
	"net/http"
	"ws-3xt/src/ws"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	port := ":8000"
	server := socketio.NewServer(nil)

	ws := ws.NewWS(server)
	ws.Run()

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Printf("Server running on %s", port)
	http.ListenAndServe(port, nil)
}

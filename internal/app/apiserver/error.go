package apiserver

import (
	"bomberman-dom/internal/app/model"

	"github.com/gorilla/websocket"
)

func (s *Server) SendError(err model.Error, client *websocket.Conn) {
	response := model.Response{Type: "error", Error: err}
	client.WriteJSON(response)
}

package apiserver

import (
	"bomberman-dom/internal/app/model"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func (s *Server) readMessage(client *websocket.Conn) {
	for {

		_, payload, err := client.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		var request model.Request

		err = json.Unmarshal(payload, &request)

		if err != nil {
			log.Println("Unmarshal payload err", err)
			return
		}

		switch request.Type {

		case "joinPlayer":
			err, player := s.AddPlayer(request.Username, client)

			if err != (model.Error{}) {
				s.SendError(err, client)
				return
			}

			s.AuthorizedJoin(player, client)
			s.BroadcastJoinedPlayer()

		case "updatePlayerPosition":
			s.UpdatePlayerPosition(client, request.Coordinates)

		case "addedBomb":
			s.HandleAddedBomb(client)

		case "sendMessage":
			s.BroadcastMessage(client, request.Message)
		}
	}

}

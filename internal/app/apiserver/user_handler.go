package apiserver

import (
	"bomberman-dom/internal/app/model"
	"time"

	"github.com/gorilla/websocket"
)

func (s *Server) AddPlayer(username string, client *websocket.Conn) (model.Error, model.Player) {
	if s.isGameStarted {
		return model.Error{Type: "gameStarted", Error: "Game already started!"}, model.Player{}
	}

	if username == "" {
		return model.Error{Type: "username", Error: "Username cannot be empty"}, model.Player{}
	}
	if len(s.players) == 4 {
		return model.Error{Type: "full", Error: "No spaces left to join the game!"}, model.Player{}
	}

	for _, p := range s.players {
		if p.Connection == client {
			return model.Error{Type: "alreadyJoined", Error: "Player already joined"}, model.Player{}
		}
	}
	coordinates := [][]int{
		{1, 0},
		{8, 9},
		{1, 9},
		{8, 0},
	}

	player := model.Player{
		Connection:     client,
		ID:             len(s.players) + 1,
		Username:       username,
		Lives:          3,
		Speed:          1,
		X:              coordinates[len(s.players)][0],
		Y:              coordinates[len(s.players)][1],
		CurrentFrame:   256,
		BombsCanSpawn:  1,
		ExplotionRange: 1,
	}

	s.players = append(s.players, player)
	return model.Error{}, player
}

func (s *Server) AuthorizedJoin(player model.Player, client *websocket.Conn) {
	response := model.Response{Type: "joinedPlayer", Player: player, Players: s.players}
	client.WriteJSON(response)
}

func (s *Server) UpdatePlayerPosition(client *websocket.Conn, coords model.Coordinates) {

	for idx, player := range s.players {
		if player.Connection == client {

			// Delay for movement
			currentTime := time.Now()
			defaultInterval := 500 * time.Millisecond
			minInterval := defaultInterval / time.Duration(player.Speed)

			if currentTime.Sub(player.LastMoveTime) < minInterval {
				return
			}

			s.players[idx].LastMoveTime = currentTime

			if frame := s.SetCurrentFrame(&player, coords); frame != 0 {
				s.players[idx].CurrentFrame = frame
			}

			if s.CheckedCollisions(player, coords) {
				s.players[idx].X = player.X + coords.X
				s.players[idx].Y = player.Y + coords.Y
			}

			s.BroadcastPlayerPosition(s.players[idx])
			break
		}

	}

}

func (s *Server) SetCurrentFrame(player *model.Player, coords model.Coordinates) int {
	if (player.X + coords.X) < player.X {
		return 64
	}
	if (player.X + coords.X) > player.X {
		return 256
	}
	if (player.Y + coords.Y) < player.Y {
		return 192
	}
	if (player.Y + coords.Y) > player.Y {
		return 128
	}
	return 0
}

func (s *Server) CheckedCollisions(player model.Player, coords model.Coordinates) bool {
	newCoordX := player.X + coords.X
	newCoordY := player.Y + coords.Y
	if newCoordX < 0 || newCoordY < 0 || newCoordX >= len(s.currentMap) || newCoordY >= len(s.currentMap[0]) {
		return false
	}

	// Check if there is another player
	for _, foundPlayer := range s.players {
		if foundPlayer.Connection != player.Connection {
			if newCoordX == foundPlayer.X && newCoordY == foundPlayer.Y {
				return false
			}
		}

	}

	// Find player index
	playerIndex := -1
	for i, foundPlayer := range s.players {
		if player.Connection == foundPlayer.Connection {
			playerIndex = i
		}
	}

	currentTile := s.currentMap[newCoordX][newCoordY]

	switch currentTile {
	case "empty_space":
		for _, powerUp := range s.powerUps {
			if powerUp.X == (player.X+coords.X) && powerUp.Y == (player.Y+coords.Y) {
				s.BroadcastRemovePowerUp(*powerUp)

				// Check for powerup type
				switch powerUp.Type {
				case "speed":
					s.players[playerIndex].Speed += 1
				case "bomb":
					s.players[playerIndex].BombsCanSpawn += 1
				case "explosion":
					s.players[playerIndex].ExplotionRange += 1
				}

			}
		}
		return true
	case "unbreakable_wall", "breakable_wall":
		return false
	}

	return true
}

func (s *Server) SendGameOver(client *websocket.Conn) {
	response := model.Response{Type: "gameOver"}
	client.WriteJSON(response)
}

func (s *Server) SendWinGame(client *websocket.Conn) {
	response := model.Response{Type: "winGame"}
	client.WriteJSON(response)
}

package apiserver

import (
	"bomberman-dom/internal/app/model"
	"time"

	"github.com/gorilla/websocket"
)

func (s *Server) HandleAddedBomb(client *websocket.Conn) {
	for idx, player := range s.players {
		if player.Connection == client {
			bomb := model.Bomb{X: player.X, Y: player.Y}
			if s.CanBombBePlaced(player, bomb) {
				s.players[idx].Bombs = append(s.players[idx].Bombs, bomb)
				go s.HandleBombTimer(bomb, idx)
				s.BroadcastPlacedBomb(bomb)
			}
			return

		}
	}
}

func (s *Server) CanBombBePlaced(player model.Player, newBomb model.Bomb) bool {
	// Check how many bombs already placed
	if len(player.Bombs) == player.BombsCanSpawn {
		return false
	}

	// Check if bomb already placed in the coordinates
	for _, bomb := range player.Bombs {
		if bomb.X == newBomb.X && bomb.Y == newBomb.Y {
			return false
		}
	}

	return true
}

func (s *Server) HandleBombTimer(bomb model.Bomb, playerIndex int) {

	// Timer for waiting bomb to explode
	time.Sleep(3 * time.Second)

	if playerIndex >= len(s.players) {
		return
	}
	// Find places to explode
	explotionCoords := s.FindPlacesToExplode(s.players[playerIndex], bomb)

	// Check if players in the coords and remove them HP

	s.DamagePlayers(explotionCoords)

	// Sending info about bomb explosion
	s.BroadcastExplotion(bomb, explotionCoords)

	s.RemoveBombFromPlayer(bomb, playerIndex)
	s.RemovePlayer()
}

func (s *Server) RemoveBombFromPlayer(bomb model.Bomb, playerIndex int) {
	for i, b := range s.players[playerIndex].Bombs {
		if b.X == bomb.X && b.Y == bomb.Y {
			s.players[playerIndex].Bombs = append(s.players[playerIndex].Bombs[:i], s.players[playerIndex].Bombs[i+1:]...)
			break
		}
	}
}

func (s *Server) FindPlacesToExplode(player model.Player, bomb model.Bomb) []model.Explotion {

	explotionCoords := []model.Explotion{}
	bombPlace := model.Explotion{X: bomb.X, Y: bomb.Y}
	explotionCoords = append(explotionCoords, bombPlace)

	// Check Right
	for i := 1; i <= player.ExplotionRange; i++ {
		x := bomb.X + i
		y := bomb.Y
		if x < len(s.currentMap) && s.currentMap[x][y] != "unbreakable_wall" {
			s.currentMap[x][y] = "empty_space"
			explotionCoords = append(explotionCoords, model.Explotion{X: x, Y: y})
		} else {
			break
		}
	}

	// Check Left
	for i := 1; i <= player.ExplotionRange; i++ {
		x := bomb.X - i
		y := bomb.Y
		if x >= 0 && s.currentMap[x][y] != "unbreakable_wall" {
			s.currentMap[x][y] = "empty_space"
			explotionCoords = append(explotionCoords, model.Explotion{X: x, Y: y})
		} else {
			break
		}
	}

	// Check Up
	for i := 1; i <= player.ExplotionRange; i++ {
		x := bomb.X
		y := bomb.Y + i
		if y < len(s.currentMap[0]) && s.currentMap[x][y] != "unbreakable_wall" {
			s.currentMap[x][y] = "empty_space"
			explotionCoords = append(explotionCoords, model.Explotion{X: x, Y: y})
		} else {
			break
		}
	}

	// Check Down
	for i := 1; i <= player.ExplotionRange; i++ {
		x := bomb.X
		y := bomb.Y - i
		if y >= 0 && s.currentMap[x][y] != "unbreakable_wall" {
			s.currentMap[x][y] = "empty_space"
			explotionCoords = append(explotionCoords, model.Explotion{X: x, Y: y})
		} else {
			break
		}
	}

	return explotionCoords
}

func (s *Server) DamagePlayers(explotionCoords []model.Explotion) {
	for _, explotionCoord := range explotionCoords {
		for idx, player := range s.players {
			if player.X == explotionCoord.X && player.Y == explotionCoord.Y {
				s.players[idx].Lives -= 1
				s.players[idx].IsDamaged = true
				s.BroadcastDamagedPlayer(s.players[idx])
				s.players[idx].IsDamaged = false
				break
			}
		}
	}
}

func (s *Server) RemovePlayer() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for idx := len(s.players) - 1; idx >= 0; idx-- {
		player := s.players[idx]
		if player.Lives == 0 {
			s.SendGameOver(player.Connection)
			s.players = append(s.players[:idx], s.players[idx+1:]...)
		}
	}

	if len(s.players) == 1 {
		s.SendWinGame(s.players[0].Connection)
		s.players = s.players[:len(s.players)-1]
		/*		s.isGameStarted = false
				s.powerUps = []*model.PowerUp{}*/
	}
}

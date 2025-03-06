package apiserver

import (
	"bomberman-dom/internal/app/model"
	"sync"

	"github.com/gorilla/websocket"
)

func (s *Server) BroadcastJoinedPlayer() {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := model.Response{Type: "joinedPlayers", Players: s.players}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))

	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()

}

func (s *Server) BroadcastTimer(timeout int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := model.Response{Type: "timer", Timer: timeout}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))

	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastGame() {
	s.mu.Lock()
	defer s.mu.Unlock()

	gameMap := s.GenerateMap()

	response := model.Response{
		Type:    "startGame",
		Players: s.players,
		Map:     model.Map{Blocks: gameMap},
	}

	s.isGameStarted = true
	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastMessage(client *websocket.Conn, description string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sender := ""

	//Search for sender username by websocket connection
	for _, player := range s.players {
		if player.Connection == client {
			sender = player.Username
		}
	}

	// If someone not joined send a message
	if sender == "" {
		return
	}

	response := model.Response{
		Type: "message",
		Message: model.Message{
			Sender:      sender,
			Description: description,
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))

	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastGameTimer(timeString string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := model.Response{Type: "gametimer", GameTimer: timeString}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastPlayerPosition(player model.Player) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := model.Response{Type: "updatePlayerPosition", Player: player}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastPlacedBomb(bomb model.Bomb) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := model.Response{Type: "addedBomb", Bomb: bomb}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastExplotion(bomb model.Bomb, explotionRange []model.Explotion) {
	s.mu.Lock()
	defer s.mu.Unlock()

	powerUpsInRange := []model.PowerUp{}

	for _, explotion := range explotionRange {
		for _, powerUp := range s.powerUps {

			if powerUp.X == explotion.X && powerUp.Y == explotion.Y {
				powerUpsInRange = append(powerUpsInRange, *powerUp)
			}
		}
	}

	response := model.Response{
		Type:      "explotion",
		Bomb:      bomb,
		Explotion: explotionRange,
		Map:       model.Map{Blocks: s.currentMap},
		Players:   s.players,
		PowerUp:   powerUpsInRange,
	}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastRemovePowerUp(powerUp model.PowerUp) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for idx := len(s.powerUps) - 1; idx >= 0; idx-- {
		pw := s.powerUps[idx]
		if pw.Type == powerUp.Type {
			s.powerUps = append(s.powerUps[:idx], s.powerUps[idx+1:]...)
		}
	}

	response := model.Response{Type: "removePowerup", SinglePowerUp: powerUp}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func (s *Server) BroadcastDamagedPlayer(player model.Player) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := model.Response{Type: "damagedPlayer", Player: player}

	wg := sync.WaitGroup{}
	wg.Add(len(s.players))
	for _, p := range s.players {
		go func(p model.Player) {
			p.Connection.WriteJSON(response)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

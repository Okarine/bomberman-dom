package apiserver

import (
	"fmt"
	"time"
)

func (s *Server) UpdateTimer() {
	timeout := 20
	for {
		if len(s.players) > 1 {
			if len(s.players) == 4 && timeout > 10 {
				timeout = 10
			}
			s.BroadcastTimer(timeout)
			timeout--
		}
		time.Sleep(1 + time.Second)
		if timeout == 0 {
			s.BroadcastGame()
			break
		}
	}
}

func (s *Server) UpdateGameTimer() {
	duration := 2 * time.Minute
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for remainingTime := duration; remainingTime > 0; remainingTime -= time.Second {
		minutes := int(remainingTime.Minutes())
		seconds := int(remainingTime.Seconds()) % 60
		timeString := fmt.Sprintf("%02d:%02d", minutes, seconds)
		s.BroadcastGameTimer(timeString)
		<-ticker.C
	}
}

package apiserver

import (
	"bomberman-dom/internal/app/model"
	"math/rand"
)

func (s *Server) GenerateMap() [][]string {

	const (
		empty = "empty"
		wall  = "wall"
		block = "random"
	)
	powerUps := []string{"speed", "explosion", "bomb"}

	var gameMap = [][]string{
		{"empty_space", "empty_space", "unbreakable_wall", "random", "random", "random", "random", "random", "empty_space", "empty_space"},
		{"empty_space", "empty_space", "random", "random", "random", "random", "random", "unbreakable_wall", "empty_space", "empty_space"},
		{"random", "random", "unbreakable_wall", "random", "random", "random", "random", "random", "random", "random"},
		{"random", "random", "random", "random", "unbreakable_wall", "random", "random", "random", "random", "random"},
		{"random", "random", "random", "random", "random", "random", "random", "random", "random", "unbreakable_wall"},
		{"unbreakable_wall", "random", "random", "random", "random", "random", "random", "random", "random", "random"},
		{"random", "random", "random", "unbreakable_wall", "random", "random", "random", "random", "random", "random"},
		{"random", "random", "random", "random", "random", "random", "random", "unbreakable_wall", "random", "random"},
		{"empty_space", "empty_space", "random", "random", "random", "random", "random", "random", "empty_space", "empty_space"},
		{"empty_space", "empty_space", "unbreakable_wall", "random", "random", "random", "unbreakable_wall", "random", "empty_space", "empty_space"},
	}

	for i := 0; i < len(gameMap); i++ {
		for j := 0; j < len(gameMap[i]); j++ {
			switch gameMap[i][j] {
			case "random":
				if rand.Float64() < 0.7 {
					gameMap[i][j] = "breakable_wall"
				} else {
					gameMap[i][j] = "empty_space"
				}
			}
		}
	}

	s.currentMap = gameMap
	breakableWalls := s.findBreakableWallCoordinates()

	newCoords := shuffleCoordinates(breakableWalls)

	for i, powerUp := range powerUps {
		pModel := model.PowerUp{
			Type: powerUp,
			X:    newCoords[i][0],
			Y:    newCoords[i][1],
		}
		s.powerUps = append(s.powerUps, &pModel)
	}

	return gameMap
}

func (s *Server) findBreakableWallCoordinates() [][]int {
	breakableWallCoordinates := [][]int{}
	for i := 0; i < len(s.currentMap); i++ {
		for j := 0; j < len(s.currentMap[i]); j++ {
			if s.currentMap[i][j] == "breakable_wall" {
				breakableWallCoordinates = append(breakableWallCoordinates, []int{i, j})
			}
		}
	}
	return breakableWallCoordinates
}

func shuffleCoordinates(coordinates [][]int) [][]int {
	for i := range coordinates {
		j := rand.Intn(i + 1)
		coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
	}

	return coordinates
}

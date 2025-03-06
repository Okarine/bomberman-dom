package model

type Response struct {
	Type    string `json:"type"`
	Player  `json:"player"`
	Players []Player `json:"players"`

	Bomb      `json:"bomb"`
	Explotion []Explotion `json:"explotion"`

	Timer     int    `json:"timer"`
	GameTimer string `json:"gametimer"`

	Map `json:"map"`

	Message `json:"message"`
	Error   `json:"error"`

	PowerUp []PowerUp `json:"powerup"`

	SinglePowerUp PowerUp `json:"singlepowerup"`
}

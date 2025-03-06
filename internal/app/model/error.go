package model

type Error struct {
	Type  string `json:"errorType"`
	Error string `json:"error"`
}

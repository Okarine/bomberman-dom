package model

import (
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	Connection     *websocket.Conn
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Lives          int    `json:"lives"`
	Speed          int    `json:"speed"`
	X              int    `json:"x"`
	Y              int    `json:"y"`
	CurrentFrame   int    `json:"currentFrame"`
	Bombs          []Bomb `json:"bombs"`
	BombsCanSpawn  int    `json:"bombsCanSpawn"`
	ExplotionRange int    `json:"explotionRange"`
	LastMoveTime   time.Time
	IsDamaged      bool `json:"isDamaged"`
}

package apiserver

import (
	"bomberman-dom/internal/app/model"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	router        *http.ServeMux
	upgrader      *websocket.Upgrader
	players       []model.Player
	mu            sync.Mutex
	isGameStarted bool
	currentMap    [][]string
	powerUps      []*model.PowerUp
}

func Start(config *Config) error {

	s := NewServer(http.NewServeMux())

	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	s.router.HandleFunc("/", s.Main)
	go s.UpdateTimer()
	go s.UpdateGameTimer()

	log.Printf("Server started at port %s\n", config.Port)
	return http.ListenAndServe(config.Port, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(router *http.ServeMux) *Server {
	upgrader := &websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &Server{
		router:        router,
		upgrader:      upgrader,
		players:       []model.Player{},
		isGameStarted: false,
	}
}

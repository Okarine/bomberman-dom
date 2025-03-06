package apiserver

import (
	"html/template"
	"log"
	"net/http"
)

func (s *Server) Main(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") == "websocket" {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go s.readMessage(conn)

	} else {
		// Render the HTML template with the data
		tmpl, err := template.ParseFiles(
			"web/templates/index.html",
			"web/templates/base.tmpl",
			"web/templates/footer.tmpl")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Problem with parsing files!"))
			return
		}
		if err = tmpl.Execute(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Problem with executing templates!"))
			return
		}
	}
}

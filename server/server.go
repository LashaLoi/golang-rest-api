package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	Port   string
	Router *mux.Router
}

// NewServer ...
func (s *Server) NewServer() error {
	log.Printf("Server was started on %v", s.Port)
	if err := http.ListenAndServe(s.Port, s.Router); err != nil {
		return err
	}

	return nil
}

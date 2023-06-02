package server

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

type Server struct {
	Engine Engine
}

func New(engine Engine) *Server {
	return &Server{
		Engine: engine,
	}
}

func (s *Server) Start() {
	log.Printf("starting server on port %s\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

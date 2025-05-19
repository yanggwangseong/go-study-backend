package main

import (
	"net/http"

	"go-study-backend/handler"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)

	return &Server{mux: mux}
}

func (s *Server) Run() error {
	return http.ListenAndServe(":8080", s.mux)
}

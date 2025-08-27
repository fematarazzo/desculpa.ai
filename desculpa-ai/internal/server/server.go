// Package server provides the HTTP server setup, including
// route registration, handlers, and static file serving.
package server

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	mux := http.NewServeMux()

	registerRoutes(mux)

	return &Server{mux: mux}
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

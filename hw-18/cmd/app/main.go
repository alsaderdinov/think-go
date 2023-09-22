package main

import (
	"net/http"
	"think-go/hw-18/pkg/api"
	"think-go/hw-18/pkg/storage"
)

func main() {
	srv := New()
	srv.Run()
}

type Server struct {
	api *api.API
}

func New() *Server {
	s := storage.New()
	srv := Server{
		api: api.New(s),
	}
	return &srv
}

func (s *Server) Run() {
	http.ListenAndServe(":8080", s.api.Router)
}

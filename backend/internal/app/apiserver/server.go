package apiserver

import (
	"test_go/internal/app/handlers/exchanges"
	"test_go/internal/app/handlers/users"
	"test_go/internal/app/store"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type Server struct {
	router *httprouter.Router
	store  store.Store
}

func NewServer(store store.Store) *Server {
	s := &Server{
		router: httprouter.New(),
		store:  store,
	}

	s.conifigureServer()

	return s
}

func (s *Server) conifigureServer() {
	users.New(s.store).Register(s.router)
	exchanges.New(s.store).Register(s.router)

}

func (s *Server) ListAndServe() error {
	handler := cors.Default().Handler(s.router)
	return http.ListenAndServe(":3001", handler)
}

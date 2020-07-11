package internal

import (
	"github.com/andrelrg/go-chat/internal/handler"
	"github.com/andrelrg/go-chat/tools"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	s := Server{r}

	s.serve()

	return &s
}

func (s *Server) serve() {

	s.Router.Path("/heartbeat").HandlerFunc(tools.PreRequest(handler.Heartbeat)).Methods(http.MethodGet)

}

package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todoapp/pkg/todoapp"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(addr string, todoService TODOService) (*Server, error) {
	s := &Server{
		addr: addr,
	}
	router := mux.NewRouter()
	s.registerRoutes(router, todoService)
	s.httpServer = &http.Server{Addr: addr, Handler: router}
	return s, nil
}

type TODOService interface {
	Insert(task *todoapp.Task) error
	FutureTasks() ([]*todoapp.Task, error)
}

func (s *Server) Start() {
	fmt.Printf("starting http server on addr %v\n", s.addr)
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic("failed to start http server")
	}
}

func (s *Server) Close() error {
	return s.httpServer.Shutdown(context.TODO())
}

func (s *Server) registerRoutes(router *mux.Router, todoService TODOService) {
	router.Methods(http.MethodPost).Path("/task").Handler(newPostTaskHandler(todoService))
	router.Methods(http.MethodGet).Path("/future_tasks").Handler(newGetFutureTasksHandler(todoService))
}

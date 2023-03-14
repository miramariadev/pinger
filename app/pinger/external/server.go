package pinger

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"

	"gitlab.studionx.ru/id/pinger/app/errors"
	"gitlab.studionx.ru/id/pinger/app/pinger/presentation/controllers"
	"gitlab.studionx.ru/id/pinger/app/pinger/presentation/middlewares"
)

type Server struct {
	addr       string
	router     *mux.Router
	errorsChan chan error
}

func NewPingerServer(
	addr string,
	errorChan chan error,
) *Server {
	return &Server{
		addr:       addr,
		router:     mux.NewRouter(),
		errorsChan: errorChan,
	}
}

func (s *Server) Run(ctx context.Context) {
	s.configureRouter(controllers.NewController())
	log.Println("Starting server on", s.addr)

	go func() {
		if err := http.ListenAndServe(s.addr, s.router); err != nil {
			s.errorsChan <- errors.NewFatalError(err)
			return
		}
	}()
	<-ctx.Done()
}

func (s *Server) configureRouter(controller *controllers.Controller) {
	s.router.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	rpc := s.router.PathPrefix("/pinger").Subrouter()

	s.registerMiddleware(rpc)
	rpc.HandleFunc("/ping", controller.HandlePing)
}

func (s *Server) registerMiddleware(router *mux.Router) {
	logRequest := middlewares.NewLogRequestMiddleware()

	router.Use(logRequest.Handle)
}

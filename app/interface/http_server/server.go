package http_server

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"math_app/app/domain/service"
	"net/http"
)

type ErrorHandler func(error)

type Server interface {
	Start() error
	Stop() error
}

type server struct {
	impl         *http.Server
	errorHandler ErrorHandler
}

func New(port uint, errorHandler ErrorHandler) Server {
	return &server{
		impl: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: newRouter(errorHandler),
		},
		errorHandler: errorHandler,
	}
}

func newRouter(errorHandler ErrorHandler) http.Handler {
	router := mux.NewRouter().PathPrefix("/api/").Subrouter()

	handler := newHandler(errorHandler)
	router.HandleFunc("/add", handler.Add).Methods(http.MethodGet)
	router.HandleFunc("/sub", handler.Sub).Methods(http.MethodGet)
	router.HandleFunc("/mul", handler.Mul).Methods(http.MethodGet)
	router.HandleFunc("/div", handler.Div).Methods(http.MethodGet)

	return router
}

func (s *server) Start() error {
	go s.run()
	return nil
}

func (s *server) Stop() error {
	return s.impl.Close()
}

func (s *server) run() {
	err := s.impl.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		s.errorHandler(err)
	}
}

func translateError(err error) int {
	switch err {
	case nil:
		return 200
	case service.ZeroDivisionError:
		fallthrough
	case InvalidArgsError:
		return 400
	default:
		return 500
	}
}

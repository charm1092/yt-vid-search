package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandler
}

func NewHTTPServer(httpHandlers *HTTPHandler) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StarServer() error {
	router := mux.NewRouter()

	router.Path("/search").Methods("GET").HandlerFunc(s.httpHandlers.HandlerSearch)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
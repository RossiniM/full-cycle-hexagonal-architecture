package server

import (
	"github.com/RossiniM/full-cycle-hexagonal-architecture/adapters/web/handler"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebServer(service application.ProductServiceInterface) *WebServer {
	return &WebServer{
		Service: service,
	}
}

func (w WebServer) Server() {
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.NewProductHandlers(r, n, w.Service)
	http.Handle("/", r)
	server := http.Server{
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

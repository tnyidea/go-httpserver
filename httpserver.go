package httpserver

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type HttpServer struct {
	Port    int
	Router  *mux.Router
	Context context.Context
}

func (p *HttpServer) ListenAndServe() error {
	// Start the Server
	log.Printf("HTTP Server Listening on port %d...", p.Port)

	return http.ListenAndServe(":"+strconv.FormatInt(int64(p.Port), 10), handlers.CORS()(p.Router))
}

package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/tnyidea/httpserver/go/httpserver/endpoints"
	"github.com/tnyidea/httpserver/go/httpserver/request"
	"github.com/tnyidea/httpserver/go/httpserver/response"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func AddDefaultRouter(router *mux.Router, ctx context.Context) *mux.Router {
	// Parameter ctx not used, but included for completeness

	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "0", "get.healthCheck")

		response.WriteDefaultResponse(w, r, endpoints.HealthCheckV0(r))
	}).Methods(http.MethodGet)

	router.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "0", "get.helloWorld")

		response.WriteDefaultResponse(w, r, endpoints.HelloWorldV0(r))
	}).Methods(http.MethodGet)

	router.HandleFunc("/helloworld/{name}", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "0", "get.helloWorldByName")

		response.WriteDefaultResponse(w, r, endpoints.HelloWorldByNameV0(r))
	}).Methods(http.MethodGet)

	return router
}

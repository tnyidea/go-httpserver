package router

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/tnyidea/httpserver/go/httpserver/request"
	"github.com/tnyidea/httpserver/go/httpserver/response"
	"github.com/tnyidea/httpserver/go/sample/userservice/data/models"
	"github.com/tnyidea/httpserver/go/sample/userservice/httpserver/config"
	"github.com/tnyidea/httpserver/go/sample/userservice/httpserver/endpoints"
	"log"
	"net/http"
)

func AddApiV1UsersRouter(router *mux.Router, ctx context.Context) *mux.Router {
	db := ctx.Value(config.UserServiceContextDatabase).(models.DB)

	router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "1", "create.users")

		response.WriteDefaultResponse(w, r, endpoints.CreateUserV1(r, db))
	}).Methods(http.MethodPost)

	router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "1", "list.users")

		response.WriteDefaultResponse(w, r, endpoints.ListUsersV1(r, db))
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/users/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "1", "get.user.uuid")

		response.WriteDefaultResponse(w, r, endpoints.GetUserByUUIDV1(r, db))
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/users/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "1", "update.user.uuid")

		response.WriteDefaultResponse(w, r, endpoints.UpdateUserV1(r, db))
	}).Methods(http.MethodPut)

	router.HandleFunc("/api/v1/users/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "1", "update.user.uuid")

		response.WriteDefaultResponse(w, r, endpoints.UpdateUserWithFieldMaskV1(r, db))
	}).Methods(http.MethodPatch)

	router.HandleFunc("/api/v1/users/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		r, requestId := request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		r = request.WithDefaultRequestApiContext(r, "1", "delete.user.uuid")

		response.WriteDefaultResponse(w, r, endpoints.DeleteUserByUUIDV1(r, db))
	}).Methods(http.MethodDelete)

	return router
}

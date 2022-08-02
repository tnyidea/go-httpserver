package endpoints

import (
	"github.com/gorilla/mux"
	"github.com/tnyidea/go-httpserver/response"
	"github.com/tnyidea/typeutils"
	"log"
	"net/http"
)

func HealthCheckV0(r *http.Request) response.DefaultResponse {
	log.Println("=== Executing HealthCheck ===")
	defer log.Println("=== HealthCheck Execution Complete ===")

	return response.NewDefaultData()
}

func HelloWorldV0(r *http.Request) response.DefaultResponse {
	log.Println("=== Executing HelloWorld ===")
	defer log.Println("=== HelloWorld Execution Complete ===")

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr("Hello World"),
		},
	}
}

func HelloWorldByNameV0(r *http.Request) response.DefaultResponse {
	log.Println("=== Executing HelloWorldByName ===")
	defer log.Println("=== HelloWorldByName Execution Complete ===")

	muxVars := mux.Vars(r)
	name := muxVars["name"]

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr("Hello, " + name),
		},
	}
}

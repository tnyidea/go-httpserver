package response

import (
	"github.com/google/uuid"
	"github.com/tnyidea/go-httpserver/request"
	"github.com/tnyidea/typeutils"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
)

// DefaultResponse
// As per https://google.github.io/styleguide/jsoncstyleguide.xml,
// with some extensions
type DefaultResponse struct {
	ApiVersion *string               `json:"version,omitempty"`
	Id         *string               `json:"id,omitempty"`
	Method     *string               `json:"method,omitempty"`
	Data       *DefaultResponseData  `json:"data,omitempty"`
	Error      *DefaultResponseError `json:"error,omitempty"`
}

type DefaultResponseData struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *string `json:"details,omitempty"`

	Items   any   `json:"items,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`

	CurrentItemCount *int `json:"currentItemCount,omitempty"`
	ItemsPerPage     *int `json:"itemsPerPage,omitempty"`
	StartIndex       *int `json:"startIndex,omitempty"`
	TotalItems       *int `json:"totalItems,omitempty"`
	PageIndex        *int `json:"pageIndex,omitempty"`
	TotalPages       *int `json:"totalPages,omitempty"`
}

type DefaultResponseError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *string `json:"details,omitempty"`
}

func NewDefaultData() DefaultResponse {
	return DefaultResponse{
		Data: &DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr(http.StatusText(http.StatusOK)),
		},
	}
}

func NewDefaultError() DefaultResponse {
	return DefaultResponse{
		Error: &DefaultResponseError{
			Code:    typeutils.IntPtr(http.StatusInternalServerError),
			Message: typeutils.StringPtr(http.StatusText(http.StatusInternalServerError)),
		},
	}
}

func WriteDefaultResponse(w http.ResponseWriter, r *http.Request, response DefaultResponse) {
	// Get context
	requestId := r.Context().Value(request.DefaultRequestId).(uuid.UUID).String()
	if requestId == "" {
		r, requestId = request.WithNewDefaultRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		log.Printf("error: Must initialize RequestContextId. Overriding value to %v", requestId)
		log.Println("calling: WriteDefaultResponse(w, r, NewDefaultError())")
		debug.PrintStack()
		WriteDefaultResponse(w, r, NewDefaultError())
		return
	}
	log.SetPrefix(requestId + ": ")

	apiVersion := r.Context().Value(request.DefaultRequestApiVersion).(string)
	if apiVersion == "" {
		log.Printf("error: Must initialize RequestContextApiVersion")
		log.Println("calling: WriteDefaultResponse(w, r, NewDefaultError())")
		debug.PrintStack()
		WriteDefaultResponse(w, r, NewDefaultError())
		return
	}

	apiMethod := r.Context().Value(request.DefaultRequestApiMethod).(string)
	if apiMethod == "" {
		log.Printf("error: Must initialize RequestContextApiMethod")
		log.Println("calling: WriteDefaultResponse(w, r, NewDefaultError())")
		debug.PrintStack()
		WriteDefaultResponse(w, r, NewDefaultError())
		return
	}

	// Validate response
	var statusCode int
	if response.Data != nil && response.Error != nil ||
		response.Data == nil && response.Error == nil {
		log.Println("error: Must specify one of Response.Data or Response.Error")
		log.Println("calling: WriteDefaultResponse(w, r, NewDefaultError())")
		debug.PrintStack()
		WriteDefaultResponse(w, r, NewDefaultError())
		return
	}
	if response.Data != nil {
		if response.Data.Items != nil {
			if reflect.TypeOf(response.Data.Items).Kind() != reflect.Slice {
				log.Println("error: Response.Data.Items must be of kind Slice")
				log.Println("calling: WriteDefaultResponse(w, r, NewDefaultError())")
				debug.PrintStack()
				WriteDefaultResponse(w, r, NewDefaultError())
				return
			}
		}
		if response.Data.Code == nil {
			response.Data.Code = typeutils.IntPtr(http.StatusOK)
		}
		statusCode = *response.Data.Code
	}
	if response.Error != nil {
		if response.Error.Code == nil {
			response.Error.Code = typeutils.IntPtr(http.StatusInternalServerError)
		}
		statusCode = *response.Error.Code
	}

	// Set response data
	response.Id = typeutils.StringPtr(requestId)
	response.ApiVersion = typeutils.StringPtr(apiVersion)
	response.Method = typeutils.StringPtr(apiMethod)

	w = WithHeaderNoCache(w)
	w = WithHeaderContentType(w, HeaderContentTypeApplicationJson)
	err := WriteJsonResponse(w, r, statusCode, response)
	if err != nil {
		log.Println("error: Error calling response.WriteJsonResponse():", err)
		debug.PrintStack()
		log.Fatal("fatal: Exiting")
	}
}

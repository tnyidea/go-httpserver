package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tnyidea/go-httpserver/request"
	"github.com/tnyidea/go-httpserver/response"
	"github.com/tnyidea/go-sample-userdata/models"
	"github.com/tnyidea/go-sample-userdata/types"
	"github.com/tnyidea/typeutils"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
)

func CreateUserV1(r *http.Request, db models.DB) response.DefaultResponse {
	log.Println("=== Executing CreateUserV1 ===")
	defer log.Println("=== CreateUserV1 Execution Complete ===")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return response.DefaultResponse{
			Error: &response.DefaultResponseError{
				Code:    typeutils.IntPtr(http.StatusBadRequest),
				Message: typeutils.StringPtr(http.StatusText(http.StatusBadRequest)),
				Details: typeutils.StringPtr(err.Error()),
			},
		}
	}

	var requestBody request.DefaultRequest
	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return response.DefaultResponse{
			Error: &response.DefaultResponseError{
				Code:    typeutils.IntPtr(http.StatusBadRequest),
				Message: typeutils.StringPtr(http.StatusText(http.StatusBadRequest)),
				Details: typeutils.StringPtr(err.Error()),
			},
		}
	}
	if requestBody.Data.Items == nil {
		log.Println("info: requestBody.Data.Items is empty")
		return response.DefaultResponse{
			Data: &response.DefaultResponseData{
				Code:    typeutils.IntPtr(http.StatusNotModified),
				Message: typeutils.StringPtr(http.StatusText(http.StatusNotModified)),
				Details: typeutils.StringPtr("Properyt data.items is empty"),
			},
		}
	}
	// TODO if requestBody.Data.Items is not a slice ....
	dataItems := requestBody.Data.Items.([]any)

	var users []types.User
	for _, dataItem := range dataItems {
		var newUser types.User
		err := request.UnmarshalRequestDataItem(dataItem, &newUser)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return response.NewDefaultError()
		}

		user, err := db.CreateUser(newUser)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return response.NewDefaultError()
		}
		users = append(users, user)
	}

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusCreated),
			Message: typeutils.StringPtr(http.StatusText(http.StatusCreated)),
			Items:   users,
		},
	}
}

func ListUsersV1(r *http.Request, db models.DB) response.DefaultResponse {
	log.Println("=== Executing FindAllUsersV1 ===")
	defer log.Println("=== FindAllUsersV1 Execution Complete ===")

	users, err := db.FindAllUsers()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return response.NewDefaultError()
	}

	totalCount, err := db.Count()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return response.NewDefaultError()
	}
	itemCount := len(users)
	itemsPerPage := itemCount
	pageIndex := 1
	totalPages := 1
	startIndex := 1 + pageIndex*itemsPerPage - itemsPerPage

	statusCode := http.StatusPartialContent
	if itemCount == totalCount {
		statusCode = http.StatusOK
	}

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(statusCode),
			Message: typeutils.StringPtr(http.StatusText(statusCode)),
			Items:   users,

			CurrentItemCount: typeutils.IntPtr(itemCount),
			ItemsPerPage:     typeutils.IntPtr(itemsPerPage),
			StartIndex:       typeutils.IntPtr(startIndex),
			TotalItems:       typeutils.IntPtr(totalCount),
			PageIndex:        typeutils.IntPtr(pageIndex),
			TotalPages:       typeutils.IntPtr(totalPages),
		},
	}
}

func GetUserByUUIDV1(r *http.Request, db models.DB) response.DefaultResponse {
	log.Println("=== Executing FindUserByUUIDV1 ===")
	defer log.Println("=== FindUserByUUIDV1 Execution Complete ===")

	muxVars := mux.Vars(r)
	uuidString := muxVars["uuid"]
	log.Println("info: Requested Id:", uuidString)

	user, err := db.FindUserByUUID(uuidString)
	if err != nil {
		if err == models.ErrorNotFound {
			log.Println("error: Requested Id", uuidString, "not found")
			return response.DefaultResponse{
				Data: &response.DefaultResponseData{
					Code:    typeutils.IntPtr(http.StatusNotFound),
					Message: typeutils.StringPtr(http.StatusText(http.StatusNotFound)),
				},
			}
		}

		log.Println(err)
		debug.PrintStack()
		return response.NewDefaultError()
	}

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr(http.StatusText(http.StatusOK)),
			Items: []types.User{
				user,
			},
		},
	}
}

func UpdateUserV1(r *http.Request, db models.DB) response.DefaultResponse {
	log.Println("=== Executing UpdateUserV1 ===")
	defer log.Println("=== UpdateUserV1 Execution Complete ===")

	muxVars := mux.Vars(r)
	uuidString := muxVars["uuid"]
	log.Println("info: Requested Id:", uuidString)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return response.DefaultResponse{
			Error: &response.DefaultResponseError{
				Code:    typeutils.IntPtr(http.StatusBadRequest),
				Message: typeutils.StringPtr(http.StatusText(http.StatusBadRequest)),
				Details: typeutils.StringPtr(err.Error()),
			},
		}
	}

	var requestBody request.DefaultRequest
	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return response.DefaultResponse{
			Error: &response.DefaultResponseError{
				Code:    typeutils.IntPtr(http.StatusBadRequest),
				Message: typeutils.StringPtr(http.StatusText(http.StatusBadRequest)),
				Details: typeutils.StringPtr(err.Error()),
			},
		}
	}
	if requestBody.Data.Items == nil {
		log.Println("info: requestBody.Data.Items is empty")
		return response.DefaultResponse{
			Data: &response.DefaultResponseData{
				Code:    typeutils.IntPtr(http.StatusNotModified),
				Message: typeutils.StringPtr(http.StatusText(http.StatusNotModified)),
				Details: typeutils.StringPtr("Properyt data.items is empty"),
			},
		}
	}
	dataItems := requestBody.Data.Items.([]any)

	var users []types.User
	for _, dataItem := range dataItems {
		var updateUser types.User
		err := request.UnmarshalRequestDataItem(dataItem, &updateUser)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return response.NewDefaultError()
		}
		if updateUser.Id == "" {
			updateUser.Id = uuidString
		}

		user, err := db.UpdateUser(updateUser)
		if err != nil {
			if err == models.ErrorNotFound {
				log.Println("error: Requested Id", uuidString, "not found")
				return response.DefaultResponse{
					Data: &response.DefaultResponseData{
						Code:    typeutils.IntPtr(http.StatusNotFound),
						Message: typeutils.StringPtr(http.StatusText(http.StatusNotFound)),
					},
				}
			}

			log.Println(err)
			debug.PrintStack()
			return response.NewDefaultError()
		}
		users = append(users, user)
	}

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr(http.StatusText(http.StatusOK)),
			Items:   users,
		},
	}
}

func UpdateUserWithFieldMaskV1(r *http.Request, db models.DB) response.DefaultResponse {
	// TODO Implement this
	return response.NewDefaultData()
}

func DeleteUserByUUIDV1(r *http.Request, db models.DB) response.DefaultResponse {
	log.Println("=== Executing DeleteUserByUUIDV1 ===")
	defer log.Println("=== DeleteUserByUUIDV1 Execution Complete ===")

	muxVars := mux.Vars(r)
	uuidString := muxVars["uuid"]
	log.Println("info: Requested Id:", uuidString)

	user, err := db.DeleteUserByUUID(uuidString)
	if err != nil {
		if err == models.ErrorNotFound {
			log.Println("error: Requested Id", uuidString, "not found")
			return response.DefaultResponse{
				Data: &response.DefaultResponseData{
					Code:    typeutils.IntPtr(http.StatusNotFound),
					Message: typeutils.StringPtr(http.StatusText(http.StatusNotFound)),
				},
			}
		}

		log.Println(err)
		debug.PrintStack()
		return response.NewDefaultError()
	}

	return response.DefaultResponse{
		Data: &response.DefaultResponseData{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr(http.StatusText(http.StatusOK)),
			Items: []types.User{
				user,
			},
			Deleted: typeutils.BoolPtr(true),
		},
	}
}

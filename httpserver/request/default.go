package request

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const DefaultRequestApiVersion = "ctxRequestApiVersion"
const DefaultRequestApiMethod = "ctxRequestApiMethod"
const DefaultRequestId = "ctxRequestId"
const DefaultRequestTimestamp = "ctxRequestTimestamp"

type DefaultRequest struct {
	Data *DefaultRequestData `json:"data,omitempty"`
}

type DefaultRequestData struct {
	Id        *string     `json:"id,omitempty"`
	FieldMask []string    `json:"fieldMask,omitempty"`
	Items     interface{} `json:"items,omitempty"`
}

func UnmarshalRequestDataItem(v interface{}, p interface{}) error {
	// TODO should we start to use any?
	// TODO this is hacky.. need a better way to turn map[string]any to the type of ptr
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, p)
	if err != nil {
		return err
	}

	return nil
}

func WithNewDefaultRequestContextId(r *http.Request) (*http.Request, string) {
	uuidValue := uuid.New()
	ctx := context.WithValue(r.Context(), DefaultRequestId, uuidValue)
	ctx = context.WithValue(ctx, DefaultRequestTimestamp, time.Now())

	return r.Clone(ctx), uuidValue.String()
}

func WithDefaultRequestApiContext(r *http.Request, version string, method string) *http.Request {
	ctx := context.WithValue(r.Context(), DefaultRequestApiVersion, version)
	ctx = context.WithValue(ctx, DefaultRequestApiMethod, method)
	return r.Clone(ctx)
}

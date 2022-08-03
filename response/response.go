package response

import (
	"encoding/json"
	"net/http"
	"time"
)

const HeaderContentTypeApplicationJson = "application/json"

func WithHeaderNoCache(w http.ResponseWriter) http.ResponseWriter {
	// Set header no-cache
	w.Header().Set("Cache-Control", "must-revalidate, no-cache, no-store, no-transform, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	return w
}

func WithHeaderContentType(w http.ResponseWriter, contentType string) http.ResponseWriter {
	// Set header content-type
	w.Header().Set("Content-Type", contentType)

	return w
}

func WriteJsonResponse(w http.ResponseWriter, r *http.Request, statusCode int, v any) error {
	// r is not used, but included for consistency

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

package helper

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// ReadRequestBody decodes the JSON request body into the provided data structure.
// It logs an error if the decoding fails.
func ReadRequestBody(r *http.Request, result interface{}) {
	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		slog.Error("failed to read request body", err)
	}
}

// WriteResponseBody writes a JSON response body to the provided HTTP response writer.
// It logs an error if the encoding fails.
func WriteResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("failed to write response body", err)
	}
}

package helper

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/adityasuryadi/messenger/internal/auth/model"
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

func WriteUnauthorizedResponse(writer http.ResponseWriter, errors error) {
	response := &model.ErrorResponse{
		Status: "UNAUTHORIZED",
		Code:   401,
		Error:  errors.Error(),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("failed to write response body", slog.String("error", err.Error()))
	}
}

func WriteBadrequestResponse(writer http.ResponseWriter, errors error) {
	response := &model.ErrorResponse{
		Status: "BAD_REQUEST",
		Code:   400,
		Error:  errors.Error(),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("failed to write response body", slog.String("error", err.Error()))
	}
}

func WriteOkResponse(writer http.ResponseWriter, data interface{}) {
	response := &model.SuccessResponse{
		Status: "OK",
		Code:   200,
		Data:   data,
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("failed to write response body", slog.String("error", err.Error()))
	}
}

func WriteInternalServerErrorResponse(writer http.ResponseWriter, errors error) {
	response := &model.ErrorResponse{
		Status: "INTERNAL_SERVER_ERROR",
		Code:   500,
		Error:  errors.Error(),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("failed to write response body", slog.String("error", err.Error()))
	}
}

func WriteNotFoundResponse(writer http.ResponseWriter, errors error) {
	response := &model.ErrorResponse{
		Status: "NOT_FOUND",
		Code:   404,
		Error:  errors.Error(),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("failed to write response body", slog.String("error", err.Error()))
	}
}

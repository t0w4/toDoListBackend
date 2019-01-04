package view

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorsResponse struct {
	Status        string   `json:"status"`
	ErrorMessages []string `json:"error_messages"`
}

type errorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func RenderInternalServerError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusInternalServerError)
	enc := json.NewEncoder(w)
	enc.Encode(&errorResponse{Status: "internal server error", ErrorMessage: message})
}

func RenderBadRequest(w http.ResponseWriter, messages []string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusBadRequest)
	enc := json.NewEncoder(w)
	enc.Encode(&errorsResponse{Status: "bad request", ErrorMessages: messages})
}

func RenderNotFound(w http.ResponseWriter, tableName string, uuid string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNotFound)
	enc := json.NewEncoder(w)
	enc.Encode(&errorResponse{
		Status:       "not found",
		ErrorMessage: fmt.Sprintf("Couldn't find %s with 'uuid'=%s", tableName, uuid),
	})
}

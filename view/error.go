package view

import (
	"encoding/json"
	"net/http"
)

type internalServerErrorResponse struct {
	Status        string   `json:"status"`
	ErrorMessages []string `json:"error_messages"`
}

func RendorInternalServerError(w http.ResponseWriter, statusCode int, messages []string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	enc.Encode(&internalServerErrorResponse{Status: "internal server error", ErrorMessages: messages})
}

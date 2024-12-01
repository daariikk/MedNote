package response

import (
	"encoding/json"
	"net/http"
)

type FailureResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendFailureResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(FailureResponse{
		Status:  "failure",
		Message: message,
	})
}

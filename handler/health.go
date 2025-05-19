package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, "OK")
}

func writeJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Message: message})
}

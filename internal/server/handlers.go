package server

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Health",
		Status: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
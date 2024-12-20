package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Msg  string
	Code int
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{"Healthcheck is successful", 200}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
	log.Print(res)
}

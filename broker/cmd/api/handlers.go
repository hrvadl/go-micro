package main

import (
	"encoding/json"
	"net/http"
)

type ResponseEnvelope struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    any    `json:"data,omitempty"`
}

func (app *Broker) HandleGetBroker(w http.ResponseWriter, r *http.Request) {
	out, _ := json.Marshal(ResponseEnvelope{
		Message: "Hit the broker",
		Error:   false,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}

package handlers

import (
	"broker/pkg/jsonh"
	"net/http"
)

type BrokerHandlers interface {
	HandleGetBroker(w http.ResponseWriter, r *http.Request)
}

type Broker struct {
	rw jsonh.JSONResponseWriterReader
}

func NewBrokerHandlers(rw jsonh.JSONResponseWriterReader) BrokerHandlers {
	return &Broker{
		rw: rw,
	}
}

func (app *Broker) HandleGetBroker(w http.ResponseWriter, r *http.Request) {
	out := jsonh.ResponseEnvelope{
		Message: "Hit the broker",
		Error:   false,
	}

	app.rw.Write(w, out, http.StatusOK)
}

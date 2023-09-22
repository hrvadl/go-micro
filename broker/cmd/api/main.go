package main

import (
	"broker/pkg/handlers"
	"broker/pkg/jsonh"
	"broker/pkg/server"
	"fmt"
	"log"
)

const webPort = "80"

func main() {
	h := handlers.NewBrokerHandlers(jsonh.NewJSONRWR())
	app := server.NewBroker(h, fmt.Sprintf(":%v", webPort))

	log.Printf("starting broker service on port %v \n", webPort)
	log.Panic(app.ListenAndServe())
}

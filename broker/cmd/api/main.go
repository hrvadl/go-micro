package main

import (
	"broker/pkg/handlers"
	"broker/pkg/jsonh"
	"broker/pkg/server"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	h := handlers.NewBrokerHandlers(jsonh.NewJSONRWR())
	app := server.NewBroker(h)

	log.Printf("starting broker service on port %v \n", webPort)

	srv := &http.Server{
		Handler: app.InitRoutes(),
		Addr:    fmt.Sprintf(":%s", webPort),
	}

	log.Panic(srv.ListenAndServe())
}

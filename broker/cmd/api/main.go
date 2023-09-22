package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Broker struct{}

func main() {
	app := &Broker{}

	log.Printf("starting broker service on port %v \n", webPort)

	srv := &http.Server{
		Handler: app.routes(),
		Addr:    fmt.Sprintf(":%s", webPort),
	}

	log.Panic(srv.ListenAndServe())
}

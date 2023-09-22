package main

import (
	"auth/pkg/db"
	"auth/pkg/handlers"
	"auth/pkg/models"
	"auth/pkg/server"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const webPort = "80"

func main() {
	d := db.MustConnect(os.Getenv("PG_DSN"))
	m := models.New(d.DB)
	h := handlers.NewAuth(&m.User)
	srv := server.NewAuth(h, fmt.Sprintf(":%v", webPort))

	log.Printf("server started on port %v", webPort)
	log.Panic(srv.ListenAndServe())
}

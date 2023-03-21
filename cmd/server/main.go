package main

import (
	"log"
	"oshikawatkm/proglog/internal/server"

	"github.com/oshikawatkm/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

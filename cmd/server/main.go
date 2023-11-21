package main

import (
	"github.com/kvngho/dsrbtlog/internal/server"
	"log"
)

func main() {
	serv := server.NewHTTPServer(":8888")
	log.Fatal(serv.ListenAndServe())
}

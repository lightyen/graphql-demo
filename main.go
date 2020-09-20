package main

import (
	"app/server"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", server.NewRouter()))
}

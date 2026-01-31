package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deadblackclover/dotshell/internal/handlers"
)

var host string
var port = "8080"

func main() {
	if envHost := os.Getenv("HOST"); envHost != "" {
		host = envHost
	}

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	http.HandleFunc("/", handlers.IndexHandler)

	log.Printf("Server is running at %s:%s\n", host, port)

	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}

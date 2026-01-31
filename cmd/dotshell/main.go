package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deadblackclover/dotshell/internal/handlers"
)

var port = "8080"

func main() {
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	http.HandleFunc("/", handlers.IndexHandler)

	log.Printf("Server is running on port: %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

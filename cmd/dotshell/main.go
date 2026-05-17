// Copyright (c) 2026, DEADBLACKCLOVER.

// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deadblackclover/dotshell/internal/handlers"
	"github.com/deadblackclover/dotshell/internal/middleware"
)

var host string
var port = "8080"
var username = ""
var password = ""

func getEnv(key string, variable *string) {
	if value := os.Getenv(key); value != "" {
		*variable = value
	}
}

func main() {
	getEnv("HOST", &host)
	getEnv("PORT", &port)
	getEnv("USERNAME", &username)
	getEnv("PASSWORD", &password)

	c := middleware.Credentials{Username: username, Password: password}

	mux := http.NewServeMux()

	mux.HandleFunc("/download", handlers.DownloadHandler)
	mux.HandleFunc("/", handlers.IndexHandler)

	addr := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Server is running at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, middleware.Authentication(mux, c)))
}

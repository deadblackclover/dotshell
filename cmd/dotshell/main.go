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

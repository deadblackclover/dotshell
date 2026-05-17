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
package middleware

import (
	"log"
	"net/http"
)

type Credentials struct {
	Username string
	Password string
}

func Authentication(next http.Handler, c Credentials) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != c.Username || password != c.Password {
			log.Printf("Authentication unsuccessful: ok=%v, username=%s", ok, username)
			w.Header().Set("WWW-Authenticate", `Basic realm="dotshell"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			log.Println("Authentication success")
		}
		next.ServeHTTP(w, r)
	})
}

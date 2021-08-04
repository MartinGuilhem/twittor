package middlew

import (
	"net/http"

	"github.com/martinguilhem/twittor/bd"
)

/* ChequeoBD is the middleware that show me the DB status */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Connection lost with Data Base", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

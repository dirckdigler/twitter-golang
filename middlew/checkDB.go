package middlew

import (
	"net/http"

	"github.com/dirckdigler/twitter/bd"
)

// CheckDB basically it is a function that acts as a handrail
// to re-validate if there is connection with the database.
func CheckDB(railing http.HandlerFunc) http.HandlerFunc {
	// Anonimous function.
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Lost connection with MongoDB", 500)
			return
		}
		// Return the same function, the goal was just to validate
		// the connection with MongoDB.
		railing.ServeHTTP(w, r)
	}
}

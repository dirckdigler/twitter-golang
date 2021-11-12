package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dirckdigler/twitter/middlew"
	"github.com/dirckdigler/twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// HandlersRoute sets port and handle routes.
func HandlersRoute() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	PORT := os.Getenv("PORT")
	// Validate if the environment variable does not exist yet.
	if PORT == "" {
		PORT = "8080"
	}
	// Object cors are the permissions given to the API to make it accessible from anywhere.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

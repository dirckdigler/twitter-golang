package main

import (
	"log"

	"github.com/dirckdigler/twitter/bd"
	"github.com/dirckdigler/twitter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("NO connection to the Mongo DataBase")
		return
	}
	handlers.HandlersRoute()
}

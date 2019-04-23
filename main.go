package main

import (
 	"log"
 	"net/http"
	"github.com/gorilla/handlers"
	"address-book/addrbk"
)

func main() {
	port := "8080"


	router := addrbk.NewRouter()

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Println("In Main........")
	// Launch server with CORS validations
 	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

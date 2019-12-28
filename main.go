package main

import (
	"go-jwt-api/app"
	"go-jwt-api/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/account", controllers.GetAccount).Methods("GET")

	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contacts", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

	log.Panicln("Listning at http://localhost:8000")
}

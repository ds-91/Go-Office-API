package main

import (
	_ "github.com/gorilla/sessions"
	_ "github.com/gorilla/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
    PORT = ":8080"
)

func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")

	// GET Requests
	r.HandleFunc("/quote/all", GetAllQuotes).Methods("GET")
	r.HandleFunc("/quote/id/{id}", GetQuoteById).Methods("GET")
	r.HandleFunc("/quote/random", GetRandomQuote).Methods("GET")
	r.HandleFunc("/quote/person/{person}", GetAllQuotesByPerson).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(PORT, r))
}

func main() {
	fmt.Println("Starting server...")
	
	// Create table for User struct
	InitialMigration()

	// Handle http requests
	handleRequests()

	// Close database when server dies
	defer db.Close()
}

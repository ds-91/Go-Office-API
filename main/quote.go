package main

import (
	"strings"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	_ "fmt"
)

var db *gorm.DB
var err error

type Quote struct {
	ID int
	Person string
	QuoteText string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "../quotes.db")
	if err != nil {
		panic("Could not open database...")
	}
	db.AutoMigrate(&Quote{})
}

func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	var quotes []Quote

	db.Find(&quotes)
	if len(quotes) == 0 {
		// Nothing found, return status 204
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Found at least 1, return status 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quotes)
}

func GetQuoteById(w http.ResponseWriter, r *http.Request) {
	quoteId := strings.TrimPrefix(r.URL.Path, "/quote/id/")
	var quote []Quote
	
	db.Where("id = ?", quoteId).First(&quote)
	if len(quote) == 0 {
		// None found, return status 204
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Found by ID, return status 200 and json encode
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quote)
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	var quote []Quote

	db.Order(gorm.Expr("random()")).First(&quote)
	if len(quote) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Random quote found, return 200 and encode
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quote)
}

func GetAllQuotesByPerson(w http.ResponseWriter, r *http.Request) {
	personFromUrl := strings.TrimPrefix(r.URL.Path, "/quote/person/")
	personToLower := strings.ToLower(personFromUrl)
	person := strings.Title(personToLower)
	var quotes []Quote

	db.Where("person = ?", person).Find(&quotes)
	if len(quotes) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quotes)
}
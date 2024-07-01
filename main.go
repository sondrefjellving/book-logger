package main

import (
	"log"

	"github.com/sondrefjellving/book-logger/internal/data_types"
	"github.com/sondrefjellving/book-logger/internal/database"
)


func main() {
	dbPath := "books.json"
	db, err := database.NewDB(dbPath)
	if err != nil {
		log.Fatal("couldn't make db")
	}

	cfg := config{
		books: make([]data_types.Book, 0),
		db: db,
	}

	cfg.setupMockData()

	startCLI(&cfg)
}

// TODOs
// - add go back to menu option for add entry page
// - refactor variable names to improve readability
// - refactor add entry command into smaller functions

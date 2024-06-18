package main

import (
	"github.com/sondrefjellving/book-logger/internal/data_types"
	mockdata "github.com/sondrefjellving/book-logger/internal/mock_data"
)


func main() {
	cfg := config{
		books: make([]data_types.Book, 0),
	}

	cfg.books = mockdata.GetBooksMock()	

	startCLI(&cfg)
}

// TODOs
// - refactor input validation (e.g: add entry command)
// - make a ui file that has methods for writing menu titles
// - refactor variable names to improve readability
// - refactor add entry command into smaller functions

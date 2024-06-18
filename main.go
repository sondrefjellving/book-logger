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
// - add a "choose in range prompt maker"
// - add go back to menu option for add entry page
// - refactor variable names to improve readability
// - refactor add entry command into smaller functions

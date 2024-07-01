package main

import (
	"fmt"

	"github.com/sondrefjellving/book-logger/internal/data_types"
	"github.com/sondrefjellving/book-logger/internal/database"
	mockdata "github.com/sondrefjellving/book-logger/internal/mock_data"
)

type config struct {
	books []data_types.Book
	db *database.DB
}

func (c *config) printBooksWithProgress() {
	for i, book := range c.books {
		fmt.Printf("%d - %s (page %d/%d)\n", i+1, book.Title, book.GetCurrentPage(), book.NumPages)
	}
}

func (c *config) printTotalPagesRead() {
	totalPages := 0
	for _, book := range c.books {
		totalPages += book.GetCurrentPage()
	}
	fmt.Printf("Total pages read: %d\n", totalPages)
}

func (c *config) setupMockData() error {
	dbStruct, err := c.db.LoadDB()
	if err != nil {
		return err
	}

	dbStruct.Books = append(dbStruct.Books, mockdata.GetBooksMock()...)
	err = c.db.WriteDB(dbStruct)
	if err != nil {
		return err
	}
	return nil
}
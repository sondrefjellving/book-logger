package main

import (
	"fmt"

	"github.com/sondrefjellving/book-logger/internal/data_types"
)

type config struct {
	books []data_types.Book
}

func (c *config) printBooksWithProgress() {
	for i, book := range c.books {
		fmt.Printf("%d - %s (page %d/%d)\n", i+1, book.Title, book.GetCurrentPage(), book.NumPages)
	}
}
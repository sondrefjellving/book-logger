package main

import (
	"fmt"

	"github.com/sondrefjellving/book-logger/internal/data_types"
)

type config struct {
	books []data_types.Book
}

func (c *config) printBooks() {
	fmt.Println("YOUR BOOKS")
	for _, book := range c.books {
		fmt.Printf("- %s\n", book.Title) // TODO: add pages read here maybe??
	}
}
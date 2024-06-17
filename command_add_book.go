package main

import (
	"fmt"

	"github.com/sondrefjellving/book-logger/internal/data_types"
)

func commandAddBook(c *config) error {
	fmt.Println("Add a new book")
	fmt.Println()

	title := GetStringFromPrompt("Title")
	author := GetStringFromPrompt("Author")
	numPages := GetIntFromPrompt("Number of pages")

	c.books = append(c.books, data_types.Book{
		Title: title,
		Author: author,
		NumPages: numPages,
		Entries: make(map[string]data_types.Entry),
	})

	return nil
}
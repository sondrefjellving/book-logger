package main

import (
	"github.com/sondrefjellving/book-logger/internal/data_types"
)

func commandAddBook(c *config) error {
	PrintPageTitle("add a new book")

	title := GetStringFromPrompt("Title")
	author := GetStringFromPrompt("Author")
	numPages := GetIntFromPrompt("Number of pages")

	c.books = append(c.books, data_types.Book{
		Title: title,
		Author: author,
		NumPages: numPages,
		Entries: make([]data_types.Entry, 0, 10),
	})

	return nil
}
package main

import (
	"fmt"
)

type Book struct {
	title 		string
	author 		string
	numPages 	int
}

func commandAddBook(c *config) error {
	fmt.Println("Add a new book")
	fmt.Println()

	title := GetStringFromPrompt("Title")
	author := GetStringFromPrompt("Author")
	numPages := GetIntFromPrompt("Number of pages")

	c.books = append(c.books, Book{
		title: title,
		author: author,
		numPages: numPages,
	})

	return nil
}
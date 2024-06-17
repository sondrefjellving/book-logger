package main

import (
	"fmt"
)

type Book struct {
	title 		string
	author 		string
	numPages 	int
	entries		map[string]Entry // TODO: check what type of date I should use
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
		entries: make(map[string]Entry),
	})

	return nil
}
package main

import "fmt"

type config struct {
	books []Book
}

func (c *config) printBooks() {
	fmt.Println("YOUR BOOKS")
	for _, book := range c.books {
		fmt.Printf("- %s\n", book.title) // TODO: add pages read here maybe??
	}
}
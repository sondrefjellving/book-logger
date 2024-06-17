package main

import "fmt"

func commandViewProgress(c *config) error {
	fmt.Println("YOUR BOOKS")
	for _, book := range c.books {
		fmt.Printf("- %s\n", book.title) // TODO: add pages read here maybe??
	}
	return nil
}
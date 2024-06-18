package main

import (
	"fmt"
	"strings"
)

func commandViewProgress(c *config) error {
	fmt.Println("YOUR BOOK PROGRESS")
	fmt.Println("(pick a book to view your entries)")
	fmt.Println()
	c.printBooksWithProgress() // add progress e.g: pages 291/401
	numOptions := len(c.books)+1
	fmt.Printf("%d: Go back to main menu\n", numOptions)
	fmt.Println()

	option := 0
	for {
		option = GetIntFromPrompt("Option")
		if isValidOption(option, numOptions) {
			break
		}
	}

	if option == numOptions { // go back to main menu
		return nil
	}
	
	bookPick := c.books[option-1]
	fmt.Println()
	fmt.Printf("YOUR ENTRIES FOR %s\n", strings.ToUpper(bookPick.Title))
	bookPick.PrintEntries()
	fmt.Println()	

	return nil
}




package main

import (
	"fmt"
)

func commandViewProgress(c *config) error {
	PrintPageTitle("your book progress")
	c.printBooksWithProgress() 

	backToMenuOption := len(c.books)+1
	fmt.Printf("%d: Go back to main menu\n", backToMenuOption)
	fmt.Println()

	option := GetIntFromPromptInRange("Option", 1, backToMenuOption)

	if option == backToMenuOption {
		return nil
	}
	
	bookPick := c.books[option-1]
	fmt.Println()
	PrintPageTitle("your entries for book") // add book title here later
	bookPick.PrintEntries()
	fmt.Println()	

	return nil
}




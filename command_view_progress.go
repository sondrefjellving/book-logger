package main

import (
	"fmt"
)

func commandViewProgress(c *config) error {
	PrintPageTitle("your book progress")
	c.printBooksWithProgress() 

	backToMenuOption := len(c.books)+1
	PrintBackToMenuOption(backToMenuOption)

	option := GetIntFromPromptInRange("Option", 1, backToMenuOption)

	if option == backToMenuOption {
		return nil
	}
	
	bookPick := c.books[option-1]
	if len(bookPick.Entries) == 0 {
		fmt.Println("You have no entries on this book... yet")
		return commandViewProgress(c)
	}
	fmt.Println()
	PrintPageTitle("your entries for book") // add book title here later
	bookPick.PrintEntries()
	fmt.Println()	

	return nil
}




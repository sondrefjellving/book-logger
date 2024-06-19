package main

import "fmt"

func commandEditBook(c *config) error {
	PrintPageTitle("edit book")
	fmt.Println("Choose a book you want to edit")
	c.printBooksWithProgress()

	backToMenuOption := len(c.books)+1
	PrintBackToMenuOption(backToMenuOption)

	option := GetIntFromPromptInRange("Book choice", 1, backToMenuOption)
	if option == backToMenuOption {
		return nil
	}

	printEditBookPage(c, option-1)
	


	return nil
}

func printEditBookPage(c *config, bookIndex int) {
	chosenBook := c.books[bookIndex]
	PrintPageTitle(chosenBook.Title)
	fmt.Println("Current:")
	chosenBook.PrintBook()
	fmt.Println()

	fmt.Println("What do you want to edit?")
	


}
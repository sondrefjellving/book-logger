package main

import "fmt"

const (
	YES = 1
	CANCEL = 2
)

func commandDeleteBook(c *config) error {
	if (len(c.books) == 0) {
		fmt.Println("There are no books to delete yet..")
		return nil
	}
	PrintPageTitle("delete book")
	c.printBooksWithProgress()

	backToMenuOption := len(c.books)+1
	PrintBackToMenuOption(backToMenuOption)

	option := GetIntFromPromptInRange("Option", 1, backToMenuOption)
	if option == backToMenuOption {
		return nil
	}

	fmt.Println()

	bookIndex := option-1
	DeleteConfirmPage(bookIndex, c)

	option = GetIntFromPromptInRange("Select option", 1, 2)

	if option == CANCEL {
		return nil
	} 

	deleteBook(c, bookIndex)
	fmt.Println()
	fmt.Println("Successfully deleted book.")
	return nil
}

func DeleteConfirmPage(bookIndex int, c *config) {
	fmt.Printf("Are you sure you want to delete %s?", c.books[bookIndex].Title)
	fmt.Println()
	fmt.Println("1 - Proceed")
	fmt.Println("2 - Cancel")
	fmt.Println()
}

func deleteBook(c *config, index int) {
	c.books = append(c.books[:index], c.books[index+1:]...)
 }
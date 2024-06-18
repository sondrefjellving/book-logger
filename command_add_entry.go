package main

import (
	"errors"
	"fmt"

	"github.com/sondrefjellving/book-logger/internal/data_types"
)


func commandAddEntry(c *config) error {
	if len(c.books) == 0 {
		return errors.New("you have to add a book first")
	}
	PrintPageTitle("add entry")

	fmt.Println("Choose a book to add an entry to")
	c.printBooksWithProgress()

	optionPrompt := fmt.Sprintf("Choose (1-%v)", len(c.books))
	option := GetIntFromPromptInRange(optionPrompt, 1, len(c.books))

	currentPage := GetIntFromPrompt("Current page")
	summary := GetStringFromPrompt("Summary")

	todaysEntry := data_types.Entry{
		Date: getCurrentDate(),
		CurrentPage: currentPage,
		Summary: summary,
	}

	c.books[option-1].AddEntry(todaysEntry)

	fmt.Println()
	fmt.Println("Added todays entry!")
	return nil
}
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

	fmt.Println("Choose a book")
	c.printBooksWithProgress()

	backToMenuOption := len(c.books)+1
	PrintBackToMenuOption(backToMenuOption)

	optionPrompt := GetChooseFromRangePrompt("Choose", backToMenuOption)
	option := GetIntFromPromptInRange(optionPrompt, 1, backToMenuOption)
	if option == backToMenuOption {
		return nil
	}

	currentPage := GetIntFromPrompt("Current page")
	summary := GetStringFromPrompt("Summary")

	todaysEntry := data_types.Entry{
		Date: getCurrentDate(),
		CurrentPage: currentPage,
		Summary: summary,
	}

	c.books[option-1].AddEntry(todaysEntry)

	fmt.Println("Added todays entry!")
	return nil
}
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
	fmt.Println("ADD ENTRY")
	fmt.Println()
	fmt.Println("Pick a book")
	for i, book := range c.books {
		fmt.Printf("%v - %s\n", i+1, book.Title)
	}
	optionPrompt := fmt.Sprintf("Pick (1-%v)", len(c.books))

	option := 0
	for {
		option = GetIntFromPrompt(optionPrompt)
		if isOptionInRange(option, len(c.books)) {
			break
		}
		fmt.Println()
		fmt.Println("Option is out of range, try again")
		fmt.Println()
	}

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

func isOptionInRange(option, max int) bool {
	return 0 < option && option <= max
}
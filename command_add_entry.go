package main

import (
	"errors"
	"fmt"
)

type Entry struct {
	currentPage	int
	summary		string
}

func commandAddEntry(c *config) error {
	if len(c.books) == 0 {
		return errors.New("you have to add a book first")
	}
	fmt.Println("ADD ENTRY")
	fmt.Println()
	fmt.Println("Pick a book")
	for i, book := range c.books {
		fmt.Printf("%v - %s\n", i+1, book.title)
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
	todaysEntry := Entry{
		currentPage: currentPage,
		summary: summary,
	}

	c.books[option-1].entries[getCurrentDate()] = todaysEntry

	fmt.Println()
	fmt.Println("Added todays entry!")
	return nil
}

func isOptionInRange(option, max int) bool {
	return 0 < option && option <= max
}
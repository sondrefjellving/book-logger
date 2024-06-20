package main

import (
	"fmt"

	"github.com/sondrefjellving/book-logger/internal/data_types"
)

const (
	EDIT_TITLE = 1
	EDIT_AUTHOR = 2
	EDIT_NUM_PAGES = 3
)

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
	chosenBook := &c.books[bookIndex]
	PrintPageTitle(chosenBook.Title)
	fmt.Println("Current:")
	chosenBook.PrintBook()
	fmt.Println()

	fmt.Println("What do you want to edit?")
	fmt.Println("1 - title")
	fmt.Println("2 - author")
	fmt.Println("3 - number of pages")
	backToMenuOption := 4
	PrintBackToMenuOption(backToMenuOption)

	option := GetIntFromPromptInRange("Chose what to edit", 1, backToMenuOption)
	if option == backToMenuOption {
		return
	}

	switch option {
		case EDIT_AUTHOR: 
			editAuthor(chosenBook)
		case EDIT_TITLE:
			editTitle(chosenBook)
		case EDIT_NUM_PAGES:
			editNumPages(chosenBook)
		default:
			return
	}

	printEditBookPage(c, bookIndex)
}

func editAuthor(b *data_types.Book) {
	fmt.Println("Current author: " + b.Author)
	author := GetStringFromPrompt("Change to")
	if err := b.EditAuthor(author); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Author updated successfully")
}
func editTitle(b *data_types.Book) {
	fmt.Println("Current title: " + b.Title)
	title := GetStringFromPrompt("Change to")
	if err := b.EditTitle(title); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Title updated successfully")
}
func editNumPages(b *data_types.Book) {
	fmt.Printf("Current number of pages: %d\n", b.NumPages)
	numPages := GetIntFromPrompt("Change to")
	if err := b.EditNumPages(numPages); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Number of pages updated successfully")
}
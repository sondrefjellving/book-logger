package data_types

import (
	"errors"
	"fmt"
	"strings"
)


type Book struct {
	Title 		string
	Author 		string
	NumPages 	int
	Entries		[]Entry 
}

type Entry struct {
	Date		string
	CurrentPage	int
	Summary		string
}


func (b *Book) EditAuthor(author string) error {
	if strings.Trim(author, " ") == "" {
		return errors.New("author cannot be blank")
	}
	b.Author = author
	return nil
}

func (b *Book) EditTitle(title string) error {
	if strings.Trim(title, " ") == "" {
		return errors.New("title cannot be blank")
	}
	b.Title = title 
	return nil
}

func (b *Book) EditNumPages(numPages int) error {
	if numPages < b.GetCurrentPage() {
		return errors.New("number of pages cannot be lower than the page you are at")
	}
	b.NumPages = numPages 
	return nil
}

func (b *Book) AddEntry(entry Entry) {
	b.Entries = append(b.Entries, entry)
}

func (b *Book) PrintBook() {
	fmt.Printf(
		"Title: %s\nAuthor: %s\nPages: %d\n",
		b.Title,
		b.Author,
		b.NumPages,
	)
}

func (b Book) GetCurrentPage() int {
	currentPage := 0
	for _, entry := range b.Entries {
		if entry.CurrentPage > currentPage {
			currentPage = entry.CurrentPage
		}
	}
	return currentPage
}

func (b Book) PrintEntries() {
	startPage := 0
	for _, entry := range b.Entries {
		fmt.Printf(
			"[%s] p. %d-%d:\nSummary: %s\n\n", 
			entry.Date, 
			startPage, 
			entry.CurrentPage, 
			entry.Summary,
		)
		startPage = entry.CurrentPage
	}
}
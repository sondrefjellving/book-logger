package data_types

import "fmt"


type Book struct {
	Title 		string
	Author 		string
	NumPages 	int
	Entries		map[string]Entry 
}

type Entry struct {
	CurrentPage	int
	Summary		string
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
	for date, entry := range b.Entries {
		fmt.Printf("[%s]: %s\n", date, entry.Summary)
	}
}
package data_types

import "fmt"


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


func (b *Book) AddEntry(entry Entry) {
	b.Entries = append(b.Entries, entry)
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
	for _, entry := range b.Entries {
		fmt.Printf("[%s], p.%d: %s\n", entry.Date, entry.CurrentPage, entry.Summary)
	}
}
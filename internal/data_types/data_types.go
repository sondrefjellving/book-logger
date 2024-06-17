package data_types


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
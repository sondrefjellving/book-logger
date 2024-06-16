package main

type Command struct {
	name 			string
	callback 		func() (error)
}

func GetCommands() map[int]Command {
	return map[int]Command{
		1: {
			name: "add entry",
			callback: commandAddEntry,
		},
		2: {
			name: "add book",
			callback: commandAddBook,
		},
	}
}
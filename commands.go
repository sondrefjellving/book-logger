package main

type Command struct {
	name 			string
	callback 		func(cfg *config) (error)
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
		3: {
			name: "edit book",
			callback: commandEditBook,
		},
		4: {
			name: "view progress",
			callback: commandViewProgress,
		},
		5: {
			name: "quit",
			callback: commandQuit,
		},
	}
}
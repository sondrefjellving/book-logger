package main

import (
	"fmt"
)

func startCLI(cfg *config) {
	commands := GetCommands()
	welcomeUser()
	for {
		showMenu(commands)
		pickOption(commands, cfg)
	}
}

func showMenu(commands map[int]Command) {
	fmt.Println("MAIN MENU")
	numMenuCommands := len(commands)
	for i := 1; i <= numMenuCommands; i++ {
		fmt.Printf("%v - %s\n", i, commands[i].name)
	}
}

func pickOption(commands map[int]Command, cfg *config) {
	fmt.Println()

	optionPrompt := fmt.Sprintf("Choose option (%v-%v)", 1, len(commands))
	option := GetIntFromPrompt(optionPrompt)
	fmt.Println()
	if isValidOption(option, len(commands)) {
		fmt.Println("Please type a number in the correct range")
		fmt.Println()
		return
	}

	err := commands[option].callback(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
}

func welcomeUser() {
	fmt.Println("Welcome, Sondre Fjellving")
	fmt.Println()
}

func isValidOption(option, numCommands int) bool {
	return option < 1 || option > numCommands
}
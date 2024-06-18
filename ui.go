package main

import (
	"fmt"
	"strings"
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
	PrintPageTitle("main menu")
	numMenuCommands := len(commands)
	for i := 1; i <= numMenuCommands; i++ {
		fmt.Printf("%v - %s\n", i, commands[i].name)
	}
}

func pickOption(commands map[int]Command, cfg *config) {
	fmt.Println()

	optionPrompt := fmt.Sprintf("Choose option (%v-%v)", 1, len(commands))
	option := GetIntFromPromptInRange(optionPrompt, 1, len(commands))
	fmt.Println()

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

func PrintPageTitle(title string) {
	fmt.Println(strings.ToUpper(title))
	fmt.Println()
}
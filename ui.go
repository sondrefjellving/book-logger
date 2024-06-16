package main

import "fmt"

func startUI() {
	commands := GetCommands()
	fmt.Println("Welcome, Sondre Fjellving")
	fmt.Println()
	showMenu(commands)
}

func showMenu(commands map[int]Command) {
	numMenuCommands := len(commands)
	for i := 1; i <= numMenuCommands; i++ {
		fmt.Printf("%v - %s\n", i, commands[i].name)
	}
}
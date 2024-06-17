package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func startUI() {
	commands := GetCommands()
	welcomeUser()
	for {
		showMenu(commands)
		pickOption(commands)
	}
}

func showMenu(commands map[int]Command) {
	numMenuCommands := len(commands)
	for i := 1; i <= numMenuCommands; i++ {
		fmt.Printf("%v - %s\n", i, commands[i].name)
	}
}

func pickOption(commands map[int]Command) {
	fmt.Print("Choose option: ")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	input := reader.Text()
	fmt.Println()

	option, err := getIntOption(input)
	if isValidOption(err, option, len(commands)) {
		fmt.Println("Please type a number in the correct range")
		fmt.Println()
		return
	}

	fmt.Println()
	commands[option].callback()
	fmt.Println()
}

func welcomeUser() {
	fmt.Println("Welcome, Sondre Fjellving")
	fmt.Println()
}

func getIntOption(input string) (int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func isValidOption(failedConversion error, option, numCommands int) bool {
	return failedConversion != nil || option < 1 || option > numCommands
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetIntFromString(input string) (int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func GetStringFromPrompt(prompt string) string {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt + ": ")
		reader.Scan()
		input := reader.Text()
		if strings.Trim(input, " ") != "" {
			return input
		}
		fmt.Println()
		fmt.Println("Invalid input, try again..")
		fmt.Println()
	}
}

func GetIntFromPrompt(prompt string) int {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt + ": ")
		reader.Scan()
		input := reader.Text()
		number, err := GetIntFromString(input)
		if strings.Trim(input, " ") != "" && err == nil {
			return number 
		}
		fmt.Println()
		fmt.Println("Invalid input, try again..")
		fmt.Println()
	}
}
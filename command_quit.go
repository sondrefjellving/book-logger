package main

import (
	"fmt"
	"os"
)

func commandQuit(c *config) error {
	fmt.Println("Shutting down..")
	fmt.Println()
	os.Exit(0)
	return nil
}
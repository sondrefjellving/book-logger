package main

import "os"

func commandQuit(c *config) error {
	os.Exit(0)
	return nil
}
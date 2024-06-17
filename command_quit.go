package main

import "os"

func commandQuit() error {
	os.Exit(0)
	return nil
}
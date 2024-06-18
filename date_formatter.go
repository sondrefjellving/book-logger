package main

import (
	"time"
)

func getCurrentDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-2-3")
}
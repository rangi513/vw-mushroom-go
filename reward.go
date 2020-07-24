package main

import (
	"fmt"
	"log"
	"os"
)

// Reward Functions have moved to individual data files mushrooms.go shuttle.go

// ScoredString : Writes scored records action:cost:probability | features to a string
func ScoredString(action int, cost float64, prob float64, features string) string {
	a := fmt.Sprint(action)
	c := fmt.Sprint(cost)
	p := fmt.Sprint(prob)
	s := a + ":" + c + ":" + p + " " + features
	return s
}

// AppendToLog : Appends scored observation to log file for replaying actions taken
// If the file doesn't exist, create it, or append to the file
func AppendToLog(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Cannot Open file", err)
	}
	if _, err := f.WriteString(text); err != nil {
		log.Fatal("Cannot write to file ", err)
	}
	if err := f.Close(); err != nil {
		log.Fatal("Cannot close file ", err)
	}
}

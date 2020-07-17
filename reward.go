package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

// GetReward : reward +5 for edible and +5, -35 with equal probability for poisonous.
// 0 reward if you don't eat
func GetReward(action int, class string) float64 {
	var reward float64
	if action == 2 && class == "e" {
		reward = 5
	} else if action == 2 && class == "p" {
		if rand.Float64() >= 0.5 {
			reward = -35
		} else {
			reward = 5
		}
	} else if action == 1 {
		reward = 0.0
	} else {
		log.Fatal("Invalid action provided.")
		fmt.Println("Action: ", action)
	}
	return reward
}

// WriteScored : Writes scored records action:cost:probability | features to a file
// currently only writes one line. If this is a batch simulation this function will need to be adjusted
func WriteScored(action int, cost float64, prob float64, features string, filepath string, logPath string) {
	a := fmt.Sprint(action)
	c := fmt.Sprint(cost)
	p := fmt.Sprint(prob)
	s := a + ":" + c + ":" + p + " " + features
	WriteToFile(filepath, s)
	appendToLog(logPath, s+"\n")
}

// appendToLog : Appends scored observation to log file for replaying actions taken
// If the file doesn't exist, create it, or append to the file
func appendToLog(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

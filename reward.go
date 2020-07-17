package main

import (
	"fmt"
	"log"
)

// GetReward : cost 20 if poisonous, cost -1 if edible, 0 if do nothing
func GetReward(action int, class string) float64 {
	var reward float64
	if action == 2 && class == "e" {
		reward = 1
	} else if action == 2 && class == "p" {
		reward = -20
	} else if action == 1 {
		reward = 0
	} else {
		log.Fatal("Invalid action provided.")
		fmt.Println("Action: ", action)
	}
	return reward
}

// WriteScored : Writes scored records action:cost:probability | features to a file
// currently only writes one line. If this is a batch simulation this function will need to be adjusted
func WriteScored(action int, cost float64, prob float64, features string, filepath string) {
	a := fmt.Sprint(action)
	c := fmt.Sprint(cost)
	p := fmt.Sprint(prob)
	s := a + ":" + c + ":" + p + " " + features
	WriteToFile(filepath, s)
}

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

// Ball : struct defining how to parse the balls csv. All strings with the class last.
type Ball struct {
	Color string
	Class string
}

// Balls : slice of Ball
type Balls []Ball

// Sample : take a single random sample from a Ball record
func (b Balls) Sample() Record {
	return b[rand.Intn(len(b))]
}

// Features : Extract features from Ball struct into VW string input
func (b Ball) Features() string {
	st := fmt.Sprintf("%+v", b)
	st = strings.TrimPrefix(st, "{")
	reg := regexp.MustCompile(`Class.*$`)
	st = reg.ReplaceAllString(st, "${1}")
	st = strings.TrimSpace(st)
	st = strings.ReplaceAll(st, ":", "=")
	st = "| " + st
	return st
}

// Reward : There are k = 2 possible states, and if the agent selects the right
// state, then reward 1 is generated. Otherwise, the agent obtains no reward (r = 0).
func (b Ball) Reward(action int) (float64, error) {
	r := 0.0
	if action == 1 && b.Class == "a" {
		r = 1.0
	} else if action == 2 && b.Class == "b" {
		r = 1.0
	}
	return r, nil
}

// GetBallActions : Get the total number of actions for the Mushroom Dataset
func GetBallActions() string {
	return "2"
}

// GetBalls : Parses the provided full csv into a Balls struct
func GetBalls() Balls {
	// Open the file
	csvFile, err := os.Open("data/balls.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(bufio.NewReader(csvFile))
	// Iterate through the records
	var bs Balls
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		bs = append(bs, Ball{
			Color: line[0],
			Class: line[1],
		})
	}
	return bs
}

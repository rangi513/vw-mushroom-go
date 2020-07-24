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
	"strconv"
	"strings"
)

// Ball : struct defining how to parse the balls csv. All strings with the class last.
type Ball struct {
	red    int
	green  int
	blue   int
	yellow int
	white  int
	black  int
	Class  string
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
	// st = strings.ReplaceAll(st, ":", "=")
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
	} else if action == 3 && b.Class == "c" {
		r = 1.0
	} else if action == 4 && b.Class == "d" {
		r = 1.0
	} else if action == 5 && b.Class == "e" {
		r = 1.0
	} else if action == 6 && b.Class == "f" {
		r = 1.0
	} else if action == 7 && b.Class == "g" {
		r = 1.0
	} else if action == 8 && b.Class == "h" {
		r = 1.0
	} else if action == 9 && b.Class == "i" {
		r = 1.0
	}
	return r, nil
}

// GetBallActions : Get the total number of actions for the Mushroom Dataset
func GetBallActions() string {
	return "6"
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
		totalFeatures := 6
		xi := make([]int, totalFeatures)
		for i, v := range line {
			if i < totalFeatures {
				vi, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal("Cannot convert str to int in csv parsing.", err)
				}
				xi[i] = vi
			}
		}
		bs = append(bs, Ball{
			red:    xi[0],
			green:  xi[1],
			blue:   xi[2],
			yellow: xi[3],
			white:  xi[4],
			black:  xi[5],
			Class:  line[6],
		})
	}
	return bs
}

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

// Shuttle :
type Shuttle struct {
	a     int
	b     int
	c     int
	d     int
	e     int
	f     int
	g     int
	h     int
	i     int
	Class int
}

// Shuttles : slice of Shutle
type Shuttles []Shuttle

// Sample :
func (s Shuttles) Sample() Record {
	return s[rand.Intn(len(s))]
}

// Features :
func (s Shuttle) Features() string {
	st := fmt.Sprintf("%+v", s)
	st = strings.TrimPrefix(st, "{")
	reg := regexp.MustCompile(`Class.*$`)
	st = reg.ReplaceAllString(st, "${1}")
	st = strings.TrimSpace(st)
	st = "| " + st
	return st
}

// Reward : . There are k = 7 possible states, and if the agent selects the right
// state, then reward 1 is generated. Otherwise, the agent obtains no reward (r = 0).
func (s Shuttle) Reward(action int) (float64, error) {
	r := 0.0
	if action == s.Class {
		r = 1.0
	}
	return r, nil
}

// GetShuttle :
func GetShuttle() Shuttles {
	// Open the file
	csvFile, err := os.Open("data/shuttle.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(bufio.NewReader(csvFile))
	// Iterate through the records
	var ss Shuttles
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		xi := make([]int, 10)
		for i, v := range line {
			vi, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Cannot convert str to int in csv parsing.", err)
			}
			xi[i] = vi
		}
		ss = append(ss, Shuttle{
			a:     xi[0],
			b:     xi[1],
			c:     xi[2],
			d:     xi[3],
			e:     xi[4],
			f:     xi[5],
			g:     xi[6],
			h:     xi[7],
			i:     xi[8],
			Class: xi[9],
		})
	}
	return ss
}

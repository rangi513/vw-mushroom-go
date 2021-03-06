package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Mushroom : struct defining how to parse the mushroom csv. All strings with the class first.
type Mushroom struct {
	Class                 string
	CapShape              string
	CapSurface            string
	CapColor              string
	Bruises               string
	Odor                  string
	GillAttachment        string
	GillSpacing           string
	GillSize              string
	GillColor             string
	StalkShape            string
	StalkRoot             string
	StalkSurfaceAboveRing string
	StalkSurfaceBelowRing string
	StalkColorAboveRing   string
	StalkColorBelowRing   string
	VeilType              string
	VeilColor             string
	RingNumber            string
	RingType              string
	SporePrintColor       string
	Population            string
	Habitat               string
}

// Mushrooms : slice of Mushroom
type Mushrooms []Mushroom

// Sample : take a single random sample from a Mushroom record
func (m Mushrooms) Sample() Record {
	return m[rand.Intn(len(m))]
}

// Features : Extract features from Mushroom struct into VW string input
func (m Mushroom) Features() string {
	s := fmt.Sprintf("%+v", m)
	s = strings.TrimSuffix(s, "}")
	s = strings.TrimPrefix(s, "{Class:")
	s = TrimLeftChar(s)
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ":", "=")
	s = "|F " + s
	return s
}

// Reward : reward +5 for edible and +5, -35 with equal probability for poisonous.
// 0 reward if you don't eat. Regret is the second item in the return set
func (m Mushroom) Reward(action int) (reward float64, regret float64) {
	class := m.Class
	if action == 2 {
		// Action Eat
		if class == "e" {
			// If edible reward = 5 regret = 0
			reward = 5.0
		} else {
			// If poisonous
			regret = (35 - 5) / 2.0
			if rand.Float64() >= 0.5 {
				reward = -35.0
			} else {
				reward = 5.0
			}
		}
	} else {
		// Action Don't Eat
		if class == "e" {
			// If edible reget = 5, if poisonous regret = 0
			regret = 5.0
		}
	}
	return reward, regret
}

// GetMushroomActions : Get the total number of actions for the Mushroom Dataset
func GetMushroomActions() []int {
	return GetActionSet(2)
}

// GetMushrooms : Parses the provided full csv into a Mushrooms struct
func GetMushrooms() Mushrooms {
	// Open the file
	csvFile, err := os.Open("data/agaricus-lepiota.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(bufio.NewReader(csvFile))
	// Iterate through the records
	var ms []Mushroom
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		ms = append(ms, Mushroom{
			Class:                 line[0],
			CapShape:              line[1],
			CapSurface:            line[2],
			CapColor:              line[3],
			Bruises:               line[4],
			Odor:                  line[5],
			GillAttachment:        line[6],
			GillSpacing:           line[7],
			GillSize:              line[8],
			GillColor:             line[9],
			StalkShape:            line[10],
			StalkRoot:             line[11],
			StalkSurfaceAboveRing: line[12],
			StalkSurfaceBelowRing: line[13],
			StalkColorAboveRing:   line[14],
			StalkColorBelowRing:   line[15],
			VeilType:              line[16],
			VeilColor:             line[17],
			RingNumber:            line[18],
			RingType:              line[19],
			SporePrintColor:       line[20],
			Population:            line[21],
			Habitat:               line[22],
		})
	}
	return ms
}

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

type mushroom struct {
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

func getMushrooms() []mushroom {
	// Open the file
	csvFile, err := os.Open("agaricus-lepiota.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(bufio.NewReader(csvFile))
	// Iterate through the records
	var mushrooms []mushroom
	for {
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		mushrooms = append(mushrooms, mushroom{
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
	return mushrooms
}

func sampleMushroom(mushrooms []mushroom) mushroom {
	return mushrooms[rand.Intn(len(mushrooms))]
}

func mushroomToString(m mushroom) string {
	s := fmt.Sprintf("%+v", m)
	s = strings.TrimSuffix(s, "}")
	s = strings.TrimPrefix(s, "{Class:")
	s = trimLeftChar(s)
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ":", "=")
	s = "| " + s
	return s
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

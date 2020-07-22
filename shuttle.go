package main

import "strconv"

// Shuttle :
type Shuttle struct {
	a     int
	b     int
	c     int
	d     int
	f     int
	g     int
	h     int
	i     int
	j     int
	Class int
}

// GetShuttle : nothing
func GetShuttle() []Shuttle {
	return []Shuttle{}
}

// Shuttles :
type Shuttles []Shuttle

// GetData :
func (s Shuttles) GetData() Shuttles {
	return nil
}

// getReardShuttle : . There are k = 7 possible states, and if the agent selects the right
// state, then reward 1 is generated. Otherwise, the agent obtains no reward (r = 0).
func getRewardShuttle(action int, class string) float64 {
	r := 0.0
	if strconv.Itoa(action) == class {
		r = 1.0
	}
	return r
}

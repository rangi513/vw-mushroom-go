package main

// Cost : Get cost from reward
func Cost(r float64) float64 {
	c := 0.0
	if r != 0.0 {
		c = r * -1.0
	}
	return c
}

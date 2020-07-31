package main

// Data : Interface for method Sample(). A slice of multiple data records
type Data interface {
	Sample() Record
}

// Record : Interface for methods Features() and Reward(). A single row of data.
type Record interface {
	Features() string
	Reward(action int) (reward float64, regret float64)
}

// CollectData : collects the dataset from the respective source and the number of total actions
func CollectData(d string) (Data, []int) {
	switch d {
	case "mushroom":
		return GetMushrooms(), GetMushroomActions()
	case "shuttle":
		return GetShuttle(), GetShuttleActions()
	case "ball":
		return GetBalls(), GetBallActions()
	}
	return nil, []int{}
}

// GetActionSet : Creates a slice of integers from 1 to N the total number of actions
func GetActionSet(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = 1 + i
	}
	return a
}

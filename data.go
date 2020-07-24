package main

// Data : Interface for method Sample(). A slice of multiple data records
type Data interface {
	Sample() Record
}

// Record : Interface for methods Features() and Reward(). A single row of data.
type Record interface {
	Features() string
	Reward(action int) float64
}

// CollectData : collects the dataset from the respective source and the number of total actions
func CollectData(d string) (Data, string) {
	switch d {
	case "mushroom":
		return GetMushrooms(), GetMushroomActions()
	case "shuttle":
		return GetShuttle(), GetShuttleActions()
	case "ball":
		return GetBalls(), GetBallActions()
	}
	return nil, "0"
}

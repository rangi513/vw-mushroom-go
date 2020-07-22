package main

type Data interface {
	Sample() Record
}

type Record interface {
	Features() string
	Reward(action int) (float64, error)
}

// CollectData : collects the dataset from the respective source
func CollectData(d string) Data {
	switch d {
	case "mushroom":
		return GetMushrooms()
		// case "shuttle":
		// return GetShuttle()
	}
	return nil
}

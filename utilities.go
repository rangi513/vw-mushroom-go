package main

import (
	"io"
	"os"
	"strconv"
)

// FileExists : checks if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

// TrimLeftChar : Trims the left character from a string
func TrimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

// GetPolicyPaths : Returns the old and new policy paths based on the current iteration
func GetPolicyPaths(policyPathBase string, iter int) (string, string) {
	op := policyPathBase + strconv.Itoa(iter) + ".vw"
	np := policyPathBase + strconv.Itoa(iter+1) + ".vw"
	return op, np
}

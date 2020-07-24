package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
)

// SelectAction : Selects action given features and policy
func SelectAction(context string, policyPath string, verbose bool) (int, float64) {
	if verbose {
		fmt.Println("Selecting Action...")
	}
	// Collect Arguments
	cmdArgs := []string{
		"-t", // testing only saves memory
		"-i", policyPath,
		"-p", "/dev/stdout", // Pipe predictions to stdout instead of file
	}
	if !verbose {
		cmdArgs = append(cmdArgs, "--quiet")
	}
	// Initialize Command
	cmd := exec.Command("vw", cmdArgs...)
	// Pipe context to stdin
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal("Could not pipe context into stdin ", err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, context)
	}()

	// Execute Command
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Couldn't select action ", err)
	}
	if verbose {
		fmt.Println("Action Selected: \n", string(stdout))
	}
	// 1. Read Action probabilities into slice
	actionProbs := getActionProbs(stdout)
	// 2. Sample From PMF
	actionIndex, probability := sampleCustomPMF(actionProbs)
	// 3. Return Action and Probability
	return actionIndex + 1, probability
}

func getActionProbs(actionTaken []byte) []float64 {
	var actionProbsFloat []float64

	actionProbs := strings.Fields(string(actionTaken))

	for _, prob := range actionProbs {
		if n, err := strconv.ParseFloat(prob, 64); err == nil {
			actionProbsFloat = append(actionProbsFloat, n)
		}
	}
	return actionProbsFloat
}

// Converted from Python Example Here
// https://vowpalwabbit.org/tutorials/cb_simulation.html#getting-a-decision-from-vowpal-wabbit
// def sample_custom_pmf(pmf):
//     total = sum(pmf)
//     scale = 1 / total
//     pmf = [x * scale for x in pmf]
//     draw = random.random()
//     sum_prob = 0.0
//     for index, prob in enumerate(pmf):
//         sum_prob += prob
//         if(sum_prob > draw):
//             return index, prob
func sampleCustomPMF(actionProbs []float64) (int, float64) {
	total := floats.Sum(actionProbs)
	scale := 1 / total
	var scaledActionProbs []float64
	for _, num := range actionProbs {
		scaledActionProbs = append(scaledActionProbs, num*scale)
	}
	draw := rand.Float64()
	sumProb := 0.0
	for index, prob := range scaledActionProbs {
		sumProb += prob
		if sumProb > draw {
			return index, prob
		}
	}
	return len(scaledActionProbs) - 1, 1.0
}

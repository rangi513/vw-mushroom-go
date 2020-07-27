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
func SelectAction(context string, policyPath string, allActions []int, verbose bool) (int, float64) {
	if verbose {
		fmt.Println("-----------------------\nSelecting Action...")
	}
	// Add Actions to Context Features (Full Context)
	fc := ScoredString(0, 0.0, 0.0, context, allActions)

	// Collect Arguments
	cmdArgs := []string{
		"-t", // testing only saves memory
		"-i", policyPath,
		"-p", "/dev/stdout", // Pipe predictions to stdout instead of file
		"--quiet", // Always need quiet because we pipe action probs to stdout
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
		io.WriteString(stdin, fc)
	}()
	if verbose {
		fmt.Println("Context Features: ", fc)
	}

	// Execute Command
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Couldn't select action ", err)
	}
	if verbose {
		fmt.Println("Action Probabilities: \n", string(stdout))
	}
	// 1. Read Action probabilities into slice
	actionIdxs, actionProbs := getActionProbs(stdout)
	// 2. Sample From PMF
	actionIndex, probability := sampleCustomPMF(actionIdxs, actionProbs)
	// 3. Return Action and Probability
	if verbose {
		fmt.Println("Action Selected: ", actionIndex)
	}
	return actionIndex, probability
}

// getActionProbs : Need to parse something like 0:0.5,2:.25,4:.25 into [0,2,4] and [0.5, 0.25, 0.25]
func getActionProbs(actionTaken []byte) ([]int, []float64) {
	ax := []int{}
	px := []float64{}
	trimmedString := strings.TrimSpace(string(actionTaken))
	sf := strings.Split(trimmedString, ",")
	for _, v := range sf {
		ss := strings.Split(v, ":")
		action, err := strconv.Atoi(ss[0])
		if err != nil {
			log.Fatal("Could not parse action probabilities", err)
		}
		prob, err := strconv.ParseFloat(ss[1], 64)
		if err != nil {
			log.Fatal("Could not parse action probabilities", err)
		}
		ax = append(ax, action+1) // +1 because index 0
		px = append(px, prob)
	}
	return ax, px
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
func sampleCustomPMF(actionIdx []int, actionProbs []float64) (int, float64) {
	total := floats.Sum(actionProbs)
	scale := 1 / total
	var scaledActionProbs []float64
	for _, num := range actionProbs {
		scaledActionProbs = append(scaledActionProbs, num*scale)
	}
	draw := rand.Float64()
	sumProb := 0.0
	for i, prob := range scaledActionProbs {
		sumProb += prob
		if sumProb > draw {
			return actionIdx[i], prob
		}
	}
	return len(scaledActionProbs) - 1, 1.0
}

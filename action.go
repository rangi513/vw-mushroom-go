package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
)

// SelectAction : Selects action given features and policy
func SelectAction(contextPath string, policyPath string, actionTakenPath string) (int, float64) {
	fmt.Println("Selecting Action...")
	cmdArgs := []string{
		"-t",
		"-d", contextPath,
		"-i", policyPath,
		"-p", actionTakenPath,
	}
	cmd := exec.Command("vw", cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Action Selected: \n", string(out))
	// 1. Read Action probabilities into slice
	actionProbs := getActionProbs(actionTakenPath)
	// 2. Sample From PMF
	actionIndex, probability := sampleCustomPMF(actionProbs)
	// 3. Return Action and Probability
	return actionIndex + 1, probability
}

func getActionProbs(actionTakenPath string) []float64 {
	var actionProbsFloat []float64

	s, err := ioutil.ReadFile(actionTakenPath)
	if err != nil {
		fmt.Print(err)
	}
	actionProbs := strings.Fields(string(s))

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
	return len(scaledActionProbs), 1.0
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set Constants and seed
	rand.Seed(time.Now().Unix())
	const iter = 5000
	// Dataset "mushroom" or "shuttle" or "ball"
	const datasetName = "ball"
	// Files
	const policyPath = "updates/policy"
	const logPath = "updates/log.dat"
	// Learning Params for CMAB Explore with Action Dependent Features
	const pEval = "dr"         // Policy Evaluation Method
	const expAlg = "--epsilon" // Exploration Algorithm
	const expParam = "0.2"     // Exploration Parameter
	// Config
	const verbose = false

	// Pull Data
	records, allActions := CollectData(datasetName)
	scoredAction := ""
	for i := 0; i <= iter-1; i++ {
		if i%500 == 0 {
			fmt.Println("Iteration: ", i)
		}
		// Define old and new policy paths for iteration
		op, np := GetPolicyPaths(policyPath, i)

		// Initialize or Update Policy
		UpdatePolicy(scoredAction, op, np, pEval, expAlg, expParam, false, verbose)

		// Sample with replacement from data
		record := records.Sample()
		featureSet := record.Features()

		// Take Action
		action, probability := SelectAction(featureSet, np, allActions, verbose)

		// Observe Reward
		reward := record.Reward(action)
		cost := Cost(reward)

		scoredAction = ScoredString(action, cost, probability, featureSet, allActions)
		AppendToLog(logPath, scoredAction+"\n")
	}

	// Final Update with coefficient output
	opp, npp := GetPolicyPaths(policyPath, iter)
	UpdatePolicy(scoredAction, opp, npp, pEval, expAlg, expParam, true, verbose)
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Set Constants and seed
	rand.Seed(time.Now().Unix())
	const iter = 10000
	// Dataset "mushroom" or "shuttle" or "ball"
	const datasetName = "shuttle"
	// Files
	const policyPath = "updates/policy"
	const logPath = "updates/log.dat"
	// Learning Params
	const banditMethod = "--cb_explore"
	const policyEvaluationApproach = "dr"
	const explorationAlgorithm = "--cover"
	const explorationParam = "2"
	// Config
	const verbose = false

	// Pull Data
	records, totalActions := CollectData(datasetName)
	scoredAction := ""
	for i := 0; i <= iter-1; i++ {
		if i%500 == 0 {
			fmt.Println("Iteration: ", i)
		}
		// Define old and new policy paths for iteration
		op, np := GetPolicyPaths(policyPath, i)

		// Initialize or Update Policy
		UpdatePolicy(scoredAction, op, np, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam, false, verbose)

		// Sample with replacement from data
		record := records.Sample()
		featureSet := record.Features()

		// Take Action
		action, probability := SelectAction(featureSet, np, verbose)
		// Observe Reward
		reward, err := record.Reward(action)
		if err != nil {
			log.Fatal("No reward returned ", err)
		}
		cost := Cost(reward)

		scoredAction = ScoredString(action, cost, probability, featureSet)
		AppendToLog(logPath, scoredAction+"\n")
	}

	// Final Update with coefficient output
	opp, npp := GetPolicyPaths(policyPath, iter)
	UpdatePolicy(scoredAction, opp, npp, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam, true, verbose)
}

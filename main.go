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
	const iter = 5000
	// Dataset "mushroom" or "shuttle"
	const datasetName = "shuttle"
	// Files
	const scoredRecordPath = "updates/scored.dat"
	const contextPath = "updates/context.dat"
	const actionTakenPath = "updates/actionTaken.dat"
	const policyPath = "updates/policy.vw"
	const logPath = "updates/log.dat"
	// Learning Params
	const banditMethod = "--cb_explore"
	const totalActions = "2"
	const policyEvaluationApproach = "dr"
	const explorationAlgorithm = "--cover"
	const explorationParam = "3"
	// Config
	const verbose = false

	// Pull Data
	records := CollectData(datasetName)
	for i := 0; i <= iter-1; i++ {
		if i%500 == 0 {
			fmt.Println("Iteration: ", i)
		}
		// Initialize or Update Policy
		UpdatePolicy(scoredRecordPath, policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam, false, verbose)

		// Sample with replacement from data
		record := records.Sample()
		featureSet := record.Features()
		WriteToFile(contextPath, featureSet)

		// Take Action
		action, probability := SelectAction(contextPath, policyPath, actionTakenPath, verbose)
		// Observe Reward
		reward, err := record.Reward(action)
		if err != nil {
			log.Fatal("No reward returned ", err)
		}
		cost := 0.0
		if reward != 0.0 {
			cost = reward * -1.0
		}
		WriteScored(action, cost, probability, featureSet, scoredRecordPath, logPath)
	}

	// Final Update with coefficient output
	UpdatePolicy(scoredRecordPath, policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam, true, verbose)
}

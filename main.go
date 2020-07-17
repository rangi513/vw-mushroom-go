package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set Constants and seed
	rand.Seed(time.Now().Unix())
	const iter = 10
	const scoredRecordPath = "data/scored.dat"
	const contextPath = "data/context.dat"
	const actionTakenPath = "data/actionTaken.dat"
	const policyPath = "data/mushroom_policy.vw"
	const logPath = "data/log.dat"
	const banditMethod = "--cb_explore"
	const totalActions = "2"
	const policyEvaluationApproach = "dr"
	const explorationAlgorithm = "--cover"
	const explorationParam = "3"
	const verbose = false

	// Initialize Policy
	UpdatePolicy(scoredRecordPath, policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam, verbose)
	// Pull Data
	mushrooms := getMushrooms()
	for i := 0; i <= iter-1; i++ {
		fmt.Println("Iteration: ", i)
		randomMushroom := sampleMushroom(mushrooms)
		featureSet := mushroomToString(randomMushroom)
		WriteToFile(contextPath, featureSet)

		// Take Action
		action, probability := SelectAction(contextPath, policyPath, actionTakenPath, verbose)
		// Observe Reward
		reward := GetReward(action, randomMushroom.Class)
		cost := 0.0
		if reward != 0.0 {
			cost = reward * -1.0
		}
		WriteScored(action, cost, probability, featureSet, scoredRecordPath, logPath)
		// Update Policy
		UpdatePolicy(scoredRecordPath, policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam, verbose)
	}
}

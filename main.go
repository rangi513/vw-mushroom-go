package main

import (
	"math/rand"
	"time"
)

func main() {
	// Set Constants and seed
	rand.Seed(time.Now().Unix())
	const iter = 2
	const scoredRecordPath = "scored.dat"
	const contextPath = "context.dat"
	const actionTakenPath = "actionTaken.dat"
	const policyPath = "mushroom_policy.vw"
	const banditMethod = "--cb_explore"
	const totalActions = "2"
	const policyEvaluationApproach = "dr"
	const explorationAlgorithm = "--cover"
	const explorationParam = "3"

	// Initialize Policy
	CreatePolicy(policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam)
	// Pull Data
	mushrooms := getMushrooms()
	for i := 0; i <= iter-1; i++ {
		randomMushroom := sampleMushroom(mushrooms)
		featureSet := mushroomToString(randomMushroom)
		WriteToFile(contextPath, featureSet)

		// Take Action
		action, probability := SelectAction(contextPath, policyPath, actionTakenPath)
		// Observe Reward
		reward := GetReward(action, randomMushroom.Class)
		cost := 0.0
		if reward != 0.0 {
			cost = reward * -1.0
		}
		WriteScored(action, cost, probability, featureSet, scoredRecordPath)
		// Update Policy
		UpdatePolicy(scoredRecordPath, policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam)
	}
}

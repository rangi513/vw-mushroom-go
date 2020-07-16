package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set Constants and seed
	rand.Seed(time.Now().Unix())
	const observedDataPath = "scored.dat"
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
	for i := 0; i <= 100; i++ {
		randomMushroom := sampleMushroom(mushrooms)

		featureSet := mushroomToString(randomMushroom)
		fmt.Println(featureSet)
		WriteToFile(contextPath, featureSet)
		// Take Action
		SelectAction(contextPath, policyPath, actionTakenPath)
		// Observe Reward
		// class := randomMushroom.Class

		// Update Policy
		// UpdatePolicy(observedDataPath, policyPath, banditMethod, totalActions, policyEvaluationApproach, explorationAlgorithm, explorationParam)
	}
}

package main

import (
	"math/rand"
	"time"

	"github.com/cheggaaa/pb"
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
	// Learning Params for CMAB Explore with Action Dependent Features
	const pEval = "dr"            // Policy Evaluation Method
	expAlg := []string{"--regcb"} // Exploration Algorithm
	// Config
	const verbose = false

	// Pull Data
	records, allActions := CollectData(datasetName)
	scoredAction := ""
	scoredActionLogs := ""
	bar := pb.StartNew(iter)
	for i := 0; i <= iter-1; i++ {
		bar.Increment()
		// Define old and new policy paths for iteration
		op, np := GetPolicyPaths(policyPath, i)

		// Initialize or Update Policy
		UpdatePolicy(scoredAction, op, np, pEval, expAlg, false, verbose)

		// Sample with replacement from data
		record := records.Sample()
		featureSet := record.Features()

		// Take Action
		action, probability := SelectAction(featureSet, np, allActions, verbose)

		// Observe Reward
		reward := record.Reward(action)
		cost := Cost(reward)

		scoredAction = ScoredString(action, cost, probability, featureSet, allActions)
		scoredActionLogs += scoredAction + "\n"
	}

	// Final Update with coefficient output
	opp, npp := GetPolicyPaths(policyPath, iter)
	UpdatePolicy(scoredAction, opp, npp, pEval, expAlg, true, verbose)
	AppendToLog(logPath, scoredActionLogs)
	bar.Finish()
}

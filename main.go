package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
)

func main() {
	// Set Constants and seed
	rand.Seed(time.Now().Unix())
	const iter = 10000
	const rounds = 5
	// Dataset "mushroom" or "shuttle" or "ball"
	datasets := []string{"ball", "shuttle", "mushroom"}
	// Files
	const policyPathBase = "updates/policy"
	const baseLogPath = "updates/log"
	// Learning Params for CMAB Explore with Action Dependent Features
	const pEval = "dr"         // Policy Evaluation Method
	expAlgSlice := [][]string{ // Exploration Algorithm
		[]string{"--regcb"},
		[]string{"--regcbopt"},
		[]string{"--cover", "1", "--nounif", "--psi", "1"},
		[]string{"--cover", "2", "--nounif", "--psi", "1"},
		[]string{"--cover", "3", "--nounif", "--psi", "1"},
		[]string{"--cover", "4", "--nounif", "--psi", "1"},
		[]string{"--cover", "8", "--nounif", "--psi", "1"},
		[]string{"--cover", "16", "--nounif", "--psi", "1"},
		[]string{"--cover", "1", "--nounif", "--psi", "0.01"},
		[]string{"--cover", "2", "--nounif", "--psi", "0.01"},
		[]string{"--cover", "3", "--nounif", "--psi", "0.01"},
		[]string{"--cover", "4", "--nounif", "--psi", "0.01"},
		[]string{"--cover", "8", "--nounif", "--psi", "0.01"},
		[]string{"--cover", "16", "--nounif", "--psi", "0.01"},
		[]string{"--cover", "1", "--nounif", "--psi", "0.1"},
		[]string{"--cover", "2", "--nounif", "--psi", "0.1"},
		[]string{"--cover", "3", "--nounif", "--psi", "0.1"},
		[]string{"--cover", "4", "--nounif", "--psi", "0.1"},
		[]string{"--cover", "8", "--nounif", "--psi", "0.1"},
		[]string{"--cover", "16", "--nounif", "--psi", "0.1"},
		[]string{"--epsilon", "0.01"},
		[]string{"--epsilon", "0.05"},
		[]string{"--epsilon", "0.1"},
		[]string{"--epsilon", "0.2"},
	}
	// Config
	const verbose = false
	for j := 0; j <= rounds-1; j++ {
		for _, datasetName := range datasets {
			// Pull Data
			records, allActions := CollectData(datasetName)
			for _, expAlg := range expAlgSlice {
				gridName := datasetName + strings.Join(expAlg, "-") + "+" + strconv.Itoa(j)
				logPath := baseLogPath + gridName + ".dat"
				policyPath := policyPathBase + gridName
				fmt.Println("Beginning: " + gridName)

				// Initalize Variables
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
		}
	}
}

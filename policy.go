package main

import (
	"fmt"
	"log"
	"os/exec"
)

// UpdatePolicy : Creates a new policy or updates an existing policy based on observed data
func UpdatePolicy(observedDataPath string, policyPath string, banditMethod string,
	totalActions string, policyEvaluationApproach string, explorationAlgorithm string,
	explorationParam string) {

	cmdArgs := []string{
		"-d", observedDataPath,
		banditMethod, totalActions,
		"--cb_type", policyEvaluationApproach,
		explorationAlgorithm, explorationParam,
		"-f", policyPath,
	}
	if fileExists(policyPath) {
		fmt.Println("Updating policy...")
		cmdArgs = append(cmdArgs, "-i", policyPath)
	} else {
		fmt.Println("Creating initial policy...")
	}
	cmd := exec.Command("vw", cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished Policy Update: \n", string(out))
}

// CreatePolicy : Creates initial policy with no data (cold start)
func CreatePolicy(initialPolicyPath string, banditMethod string,
	totalActions string, policyEvaluationApproach string, explorationAlgorithm string,
	explorationParam string) {
	fmt.Println("Creating Initial Policy")
	cmdArgs := []string{
		banditMethod, totalActions,
		"--cb_type", policyEvaluationApproach,
		explorationAlgorithm, explorationParam,
		"-f", initialPolicyPath,
	}
	cmd := exec.Command("vw", cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished Creating Initial Policy: \n", string(out))
}

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// UpdatePolicy : Creates a new policy or updates an existing policy based on observed data
func UpdatePolicy(observedDataPath string, policyPath string, banditMethod string,
	totalActions string, policyEvaluationApproach string, explorationAlgorithm string,
	explorationParam string, verbose bool) {

	cmdArgs := []string{
		banditMethod, totalActions,
		"--cb_type", policyEvaluationApproach,
		explorationAlgorithm, explorationParam,
		"-f", policyPath,
		"--invert_hash", strings.TrimRight(policyPath, ".vw") + ".txt",
	}
	if fileExists(policyPath) {
		// If policy exists, update existing policy
		fmt.Println("Updating policy...")
		cmdArgs = append(cmdArgs, "-d", observedDataPath, "-i", policyPath)
	} else {
		// If policy does not exist, create a new one
		fmt.Println("Creating initial policy...")
	}
	if !verbose {
		cmdArgs = append(cmdArgs, "--quiet")
	}
	cmd := exec.Command("vw", cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished Policy Update: \n", string(out))
}

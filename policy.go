package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

// UpdatePolicy : Creates a new policy or updates an existing policy based on observed data
func UpdatePolicy(scoredAction string, oldPolicyPath string, newPolicyPath string,
	policyEvaluationApproach string, explorationAlgorithm string,
	explorationParam string, coefOutput bool, verbose bool) {

	// Format Command Arguments
	cmdArgs := []string{
		"--cb_explore_adf",
		"--cb_type", policyEvaluationApproach,
		explorationAlgorithm, explorationParam, //"--nounif",
	}
	if FileExists(oldPolicyPath) {
		// If policy exists, update existing policy
		if verbose {
			fmt.Println("Updating policy...")
		}
		cmdArgs = append(cmdArgs, "-i", oldPolicyPath)
	} else {
		// If policy does not exist, create a new one
		if verbose {
			fmt.Println("Creating initial policy...")
		}
		cmdArgs = append(cmdArgs, "-q", "AF")
	}
	cmdArgs = append(cmdArgs,
		"--save_resume",
		"-f", newPolicyPath)
	if coefOutput {
		cmdArgs = append(cmdArgs, "--invert_hash", strings.TrimSuffix(newPolicyPath, ".vw")+".txt")
	}
	if !verbose {
		cmdArgs = append(cmdArgs, "--quiet")
	}
	cmd := exec.Command("vw", cmdArgs...)

	// Write Data in Stdin
	if scoredAction != "" {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, scoredAction)
		}()
	}

	// Execute command
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	exec.Command("rm", oldPolicyPath).CombinedOutput()
	if verbose {
		fmt.Println("Finished Policy Update: \n", string(out))
	}
}

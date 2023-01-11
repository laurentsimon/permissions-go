package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	permissions "github.com/laurentsimon/permissions-go"
)

func generatePolicyValidate(args map[string]string) map[string]string {
	msg := map[string]string{
		"--trace-file string":  "Trace file to use as input",
		"--policy-file string": "Policy file to write the policy",
	}

	if len(args) != 2 {
		return msg
	}

	_, ok := args["--trace-file"]
	if !ok {
		return msg
	}

	_, ok = args["--policy-file"]
	if !ok {
		return msg
	}
	return nil
}

func generatePolicy(args map[string]string) error {
	tracePath, _ := args["--trace-file"]
	// policyFile, _ := args["--policy-file"]
	traceFile, err := os.Open(tracePath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fileScanner := bufio.NewScanner(traceFile)
	fileScanner.Split(bufio.ScanLines)
	var traces []permissions.AccessMetadata
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var metadata permissions.AccessMetadata
		err = json.Unmarshal([]byte(line), &metadata)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		traces = append(traces, metadata)
	}

	one := 1
	pol := permissions.PolicyDefaultDisallow
	c := &permissions.Config{
		Version: &one,
		Default: &pol,
	}

	fmt.Println(*c)

	// Check what each dependency needs access to.
	for _, m := range traces {
		walkTrace(&m)
	}

	// TODO: add all remaining deps
	return nil
}

func walkTrace(m *permissions.AccessMetadata, c *permissions.Config) {
	// TODO: add to config.
	for _, e := range m.Trace {
		pkg := e.Pkg
		fmt.Println(pkg)
		// TODO: output sticky bit?
		panic("hey")
	}
}

package main

import (
	"fmt"
	"os"
)

var cli = "permission-cli"

func main() {
	args := os.Args
	if len(args) <= 1 {
		usage(cli, nil, nil)
	}

	cmd := args[1]
	m := argsAsMap(args[2:])
	switch cmd {
	case "generate-policy":
		errs := generatePolicyValidate(m)
		if errs != nil {
			usage(cli, &cmd, errs)
			os.Exit(1)
		}
		err := generatePolicy(m)
		exit(err)
	case "access":
		notImplemented()
		os.Exit(1)
	case "help":
		usage(cli, nil, nil)
	default:
		usage(cli, nil, nil)
		os.Exit(1)
	}
}

func usage(cli string, cmd *string, errs map[string]string) {
	if cmd == nil {
		s := `Usage: %s option [args]

Available Commands:
  generate-policy   Generate a policy file
  help              Help about any command
  access            Determine if the program of a dependency has access to a resource`
		fmt.Fprintln(os.Stderr, fmt.Sprintf(s, cli))
	} else {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Usage: %s %s [args]\n", cli, *cmd))
		for k, v := range errs {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s\t%s", k, v))
		}
	}
	fmt.Fprintln(os.Stderr)
}

func notImplemented() {
	fmt.Fprintln(os.Stderr, "not implemented")
	os.Exit(1)
}

func argsAsMap(args []string) map[string]string {
	m := make(map[string]string)
	i := 0
	for i < len(args)/2 {
		m[args[2*i]] = args[(2*i)+1]
		i++
	}

	return m
}

func exit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("runtime error: %s", err))
		os.Exit(1)
	}
}

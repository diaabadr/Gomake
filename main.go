package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	makefile "github.com/codescalersinternships/gomake-Diaa/internal"
)

var ErrMissingMakefileArg = fmt.Errorf("make: option requires an argument -- 'f'")

const HelpMessage = `Usage: make [options] [target] ...
Options:
  -f FILE`

func main() {
	filePath, target, err := ParseCommand()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing command: %v\n", err)
		fmt.Println(HelpMessage)
		os.Exit(1)
	}

	adjList, targToCmds, err := makefile.ReadMakefile(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	depGraph := makefile.NewDependencyGraph()

	depGraph.SetAdjacencyList(adjList)
	depGraph.SetTargetToCommands(targToCmds)

	err = depGraph.ExecuteTargetKAndItsDeps(target)

	if err != nil {
		fmt.Fprintf(os.Stderr, "execution error: %v\n", err)
		os.Exit(1)
	}

}

func ParseCommand() (string, string, error) {
	filePath := flag.String("f", "", "name of the file to be explored")

	flag.Parse()

	if len(flag.Args()) != 1 {
		return "", "", errors.New("please specify a single target")
	}
	target := flag.Args()[0]

	if *filePath == "" {
		*filePath = "Makefile"
	}

	return *filePath, target, nil
}

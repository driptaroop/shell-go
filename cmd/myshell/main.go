package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		commandInput()
	}
}

func commandInput() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	command = strings.TrimSpace(command)
	validateCommand(command)
}

func validateCommand(command string) {
	// create an array of commands
	commands := []string{"exit 0"}

	// check if string is present in an array
	if contains(commands, command) {
		evaluateCommand(command)
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func evaluateCommand(command string) {
	switch command {
	case "exit 0":
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "Command found")
	}
}

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

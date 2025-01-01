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
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	command = strings.TrimSpace(command)

	// create an array of commands
	commands := []string{}

	// check if string is present in an array
	if contains(commands, command) {
		fmt.Fprintf(os.Stdout, "Command found")
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
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

package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shellbuiltins"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	shellbuiltins.Commands = shellbuiltins.BuiltinDefinition()
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
	//get all except the first word in command
	command, argument, _ := strings.Cut(command, " ")
	arguments := PreprocessArguments(argument)

	if f, ok := shellbuiltins.Commands[command]; ok {
		evaluateCommand(f, arguments)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir, _ := os.ReadDir(path)
			for _, file := range dir {
				if file.Name() == command {
					output, _ := exec.Command(command, arguments...).Output()
					fmt.Fprintf(os.Stdout, "%s", output)
					return
				}
			}
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func evaluateCommand(f func([]string), arguments []string) {
	f(arguments)
}

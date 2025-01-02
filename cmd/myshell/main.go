package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var commands map[string]func(string)

func main() {
	commands = builtinDefinition()
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

	if f, ok := commands[command]; ok {
		evaluateCommand(f, argument)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir, _ := os.ReadDir(path)
			for _, file := range dir {
				if file.Name() == command {
					output, _ := exec.Command(command, argument).Output()
					fmt.Fprintf(os.Stdout, "%s", output)
					return
				}
			}
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func evaluateCommand(f func(string), argument string) {
	f(argument)
}

func builtinDefinition() map[string]func(string) {
	return map[string]func(string){
		"exit": func(exitCode string) {
			code, _ := strconv.Atoi(exitCode)
			os.Exit(code)
		},
		"echo": func(message string) {
			fmt.Println(message)
		},
		"type": func(command string) {
			if _, ok := commands[command]; ok {
				fmt.Printf("%s is a shell builtin\n", command)
			} else {
				paths := strings.Split(os.Getenv("PATH"), ":")
				for _, path := range paths {
					dir, _ := os.ReadDir(path)
					for _, file := range dir {
						if file.Name() == command {
							fmt.Printf("%s is %s/%s\n", command, path, command)
							return
						}
					}
				}
				fmt.Printf("%s: not found\n", command)
			}
		},
	}
}

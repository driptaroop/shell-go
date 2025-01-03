package shellbuiltins

import (
	"fmt"
	"os"
	"strings"
)

func Type(command []string) {
	if _, ok := Commands[command[0]]; ok {
		fmt.Printf("%s is a shell builtin\n", command[0])
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir, _ := os.ReadDir(path)
			for _, file := range dir {
				if file.Name() == command[0] {
					fmt.Printf("%s is %s/%s\n", command[0], path, command[0])
					return
				}
			}
		}
		fmt.Printf("%s: not found\n", command[0])
	}
}

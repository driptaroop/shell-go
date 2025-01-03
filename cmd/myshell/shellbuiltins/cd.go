package shellbuiltins

import (
	"fmt"
	"os"
	"strings"
)

func Cd(path []string) {
	currentUser, _ := os.UserHomeDir()
	err := os.Chdir(strings.ReplaceAll(path[0], "~", currentUser))
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path[0])
	}
}

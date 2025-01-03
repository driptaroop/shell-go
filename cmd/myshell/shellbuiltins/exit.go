package shellbuiltins

import (
	"os"
	"strconv"
)

func Exit(exitCode []string) {
	code, _ := strconv.Atoi(exitCode[0])
	os.Exit(code)
}

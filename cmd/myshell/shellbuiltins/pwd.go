package shellbuiltins

import (
	"fmt"
	"os"
)

func Pwd(_ []string) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

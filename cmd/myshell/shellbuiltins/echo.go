package shellbuiltins

import (
	"fmt"
	"strings"
)

func Echo(message []string) {
	fmt.Println(strings.Join(message, " "))
}

package main

import "strings"

func PreprocessArguments(arg string) []string {
	startsWithSingleQuote := strings.HasPrefix(arg, "'")
	endsWithSingleQuote := strings.HasSuffix(arg, "'")
	if startsWithSingleQuote && endsWithSingleQuote {
		arg = strings.Trim(arg, "'")
		return strings.Split(arg, "' '")
	} else {
		fields := strings.Fields(arg)
		return fields
	}
}

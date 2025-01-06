package main

import "strings"

func PreprocessArguments(arg string) []string {
	startsWithSingleQuote := strings.HasPrefix(arg, "'")
	endsWithSingleQuote := strings.HasSuffix(arg, "'")
	startsWithDDoubleQuote := strings.HasPrefix(arg, "\"")
	endsWithDoubleQuote := strings.HasSuffix(arg, "\"")
	if startsWithSingleQuote && endsWithSingleQuote {
		arg = strings.Trim(arg, "'")
		return strings.Split(arg, "' '")
	} else if startsWithDDoubleQuote && endsWithDoubleQuote {
		fields := strings.Fields(arg)
		// Remove the double quotes from the fields
		for i, field := range fields {
			fields[i] = strings.Trim(field, "\"")
		}
		return fields
	} else {
		fields := strings.Fields(arg)
		return fields
	}
}

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
		fields := processDoubleQuote(arg)
		return fields
	} else {
		fields := strings.Fields(arg)
		return fields
	}
}

func processDoubleQuote(arg string) []string {
	arg = strings.Trim(arg, "\"")
	fields := strings.Fields(arg)
	for i, field := range fields {
		if strings.Contains(field, " ") {
			fields[i] = strings.ReplaceAll(field, " ", "")
		}
	}
	return fields
}

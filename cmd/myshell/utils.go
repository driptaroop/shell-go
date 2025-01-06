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
	// traverse the string and tokenize it with double quotes as delimiter
	arg = strings.ReplaceAll(arg, "\"\"", "")
	fields := make([]string, 0)
	var field string
	var inQuote bool
	for _, r := range arg {
		if r == '"' {
			if inQuote {
				fields = append(fields, field)
				field = ""
			}
			inQuote = !inQuote
		} else {
			field += string(r)
		}
	}
	return fields
}

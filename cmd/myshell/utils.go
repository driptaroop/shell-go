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
	// travserse the string and tokenize it with double quotes as delimiter
	var field = ""
	var inQuote bool
	for i, r := range arg {
		if r == '"' {
			if inQuote {
				inQuote = false
			} else {
				inQuote = true
			}
		} else {
			if inQuote == false && r == ' ' && arg[i-1] == ' ' {
				continue
			}
			field += string(r)
		}
	}
	return []string{field}
}

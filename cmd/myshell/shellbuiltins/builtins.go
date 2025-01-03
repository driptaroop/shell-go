package shellbuiltins

var Commands map[string]func([]string)

func BuiltinDefinition() map[string]func([]string) {
	return map[string]func([]string){
		"exit": Exit,
		"echo": Echo,
		"pwd":  Pwd,
		"type": Type,
		"cd":   Cd,
	}
}

package main 

import (
	"fmt" 
	"strings"
	// "code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string] int {
	return map[string] int { "x" : 1}
}

func main() {
	words :=  strings.Fields(" Welcome to my world Welcome   ")
	fmt.Println("Fields are %q", words)
}

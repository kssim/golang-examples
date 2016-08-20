package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var result string

	// Basic example
	for i := 1; i < len(os.Args); i++ {
		result += os.Args[i] + " "
	}
	fmt.Println(result)

	// Use range example
	result = ""
	for _, arg := range os.Args[1:] {
		result += arg + " "
	}
	fmt.Println(result)

	// Use Join method example
	fmt.Println(strings.Join(os.Args[1:], " "))
}

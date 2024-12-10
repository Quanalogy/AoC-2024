package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	filename := os.Args[1]
	// filename := "input/part1_example"

	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := string(dataBytes)

	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	inputSplit := strings.Split(input, "\n\n")
	rulesInput := inputSplit[0]
	pagesInput := inputSplit[1]

	rules := GetPageOrderingRules(rulesInput)
	pages := GetPageOrdering(pagesInput)

	// fmt.Printf("rules: %v\n", rules)
	// fmt.Printf("pages: %v\n", pages)

	fmt.Printf("%v\n", GetMiddleNumberJoinedOfValidPages(pages, rules))
}

package main

import "os"

func main() {
	// if len(os.Args) != 2 {
	// 	println("Usage: go run main.go <input_file>")
	// 	os.Exit(1)
	// }
	// filename := os.Args[1]
	filename := "input/part1"

	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := string(dataBytes)

	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	part1 := Xmas(input)
	println(part1)
	part2 := Xmas_part2(input)
	println(part2)
}

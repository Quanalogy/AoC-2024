package main

import "os"

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	filename := os.Args[1]

	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := string(dataBytes)
	part1 := Multiply(input)
	part2 := Multiply_part2(input)
	println(part1)
	println(part2)
}
package main

import "os"

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	filename := os.Args[1]
	// filename := "./input/part1"

	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(dataBytes)
	reports := GetReports(data)
	safeReports := GetSafeReports(reports)
	safeReports_part2 := GetSafeReports_part2(reports)

	println("Amount of safe reports: ", len(safeReports))
	println("Amount of safe reports part 2: ", len(safeReports_part2))
}

package main

import (
	"os"
)

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
	data := string(dataBytes)

	numbers := GetNumbers(data)
	odd, even := ListifyNumbersSorted(numbers)

	totalDistance := TotalDistance(odd, even)
	totalSimilarity := TotalSimilarity(odd, even)

	println("Distance: ", totalDistance)
	println("Similarity: ", totalSimilarity)
}

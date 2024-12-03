package main

import (
	"regexp"
	"slices"
	"strconv"
)

func GetNumbers(data string) []int {
	numbersRegex := regexp.MustCompile(`\d+`)
	numbers := numbersRegex.FindAllString(data, -1)
	result := make([]int, len(numbers))

	for i, number := range numbers {
		val, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		result[i] = val
	}

	return result
}

func ListifyNumbersSorted(data []int) (odd, even []int) {
	for i := 0; i < len(data); i += 2 {
		even = append(even, data[i])
		odd = append(odd, data[i+1])
	}

	slices.Sort(odd)
	slices.Sort(even)

	return odd, even
}

func Distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func TotalDistance(a, b []int) int {
	total := 0

	for i := 0; i < len(a); i++ {
		total += Distance(a[i], b[i])
	}

	return total
}

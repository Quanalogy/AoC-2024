package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetReports(data string) [][]int {
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	numbersRegex := regexp.MustCompile(`\d+`)
	numberLines := strings.Split(data, "\n")
	result := make([][]int, len(numberLines))

	for i, numberLine := range numberLines {
		numbers := numbersRegex.FindAllString(numberLine, -1)
		result[i] = make([]int, len(numbers))

		for j, number := range numbers {
			val, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			result[i][j] = val
		}
	}

	return result
}

func ReportSafe(report []int) bool {
	last := report[0]
	increasing := report[0] < report[1]

	for i := 1; i < len(report); i++ {
		current := report[i]

		if increasing && current < last {
			return false
		}

		if !increasing && current > last {
			return false
		}

		difference := math.Abs(float64(last - current))
		if difference < 1 || difference > 3 {
			return false
		}

		last = current
	}

	return true
}

func GetSafeReports(reports [][]int) [][]int {
	var validReports [][]int

	for _, report := range reports {
		if ReportSafe(report) {
			validReports = append(validReports, report)
		}
	}

	return validReports
}

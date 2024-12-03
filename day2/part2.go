package main

import (
	"math"
	"slices"
)

func IsSafeNumber(a, b int, increasing bool) bool {
	difference := math.Abs(float64(a - b))
	return difference >= 1 && difference <= 3 && increasing == (a < b)
}

func GetDirections(report []int) (int, int) {
	increases := 0
	decreases := 0

	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			increases++
		} else {
			decreases++
		}
	}

	return increases, decreases
}

// gave up naming this function
func GetSafe_part2(report []int) ([]bool, int) {
	increases, decreases := GetDirections(report)
	reportLen := len(report)
	isSafe := make([]bool, reportLen-1)
	safeCount := 0
	increasing := increases > decreases

	for i := 1; i < reportLen; i++ {
		current := report[i]
		prev := report[i-1]
		safe := IsSafeNumber(prev, current, increasing)

		if safe {
			safeCount++
		} else {
			safeCount--
		}

		isSafe[i-1] = safe
	}

	return isSafe, safeCount
}

func ReportSafe_part2(report []int) bool {
	dampenedReport := report

	for i := 0; i < 2; i++ {
		increases, decreases := GetDirections(dampenedReport)
		if increases > 1 && decreases > 1 {
			return false
		}

		reportLen := len(dampenedReport)
		isSafe, safeCount := GetSafe_part2(dampenedReport)
		safeDistance := math.Abs(float64(safeCount)) - float64(reportLen) + 1

		if safeDistance > 1 {
			return false
		}

		if safeDistance == 0 {
			return true
		}

		indexToRemove := slices.Index(isSafe, safeCount < 0)
		dampenedReport = append(report[:indexToRemove], report[indexToRemove+1:]...)
	}

	return false
}

func GetSafeReports_part2(reports [][]int) [][]int {
	var validReports [][]int

	for _, report := range reports {
		if ReportSafe_part2(report) {
			validReports = append(validReports, report)
		}
	}

	return validReports
}

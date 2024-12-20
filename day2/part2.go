package main

import (
	"fmt"
	"math"
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
		} else if report[i] > report[i+1] {
			decreases++
		}
	}

	return increases, decreases
}

type DirectionReport int

const (
	MoreThanOneDifferingElement DirectionReport = iota
	Uniform
	OneDifferingElement
	NoDifferingElements
)

// returns -1 if it's not safe to process
// returns 0 if it's safe to process
// returns 1 if the
func GetReportDirectionReport(report []int) DirectionReport {
	increases, decreases := GetDirections(report)

	if increases == 0 && decreases == 0 {
		return NoDifferingElements
	}
	if increases > 1 && decreases > 1 {
		return MoreThanOneDifferingElement
	}
	if increases == 1 || decreases == 1 {
		return OneDifferingElement
	}
	return Uniform
}

func RemoveElementNotFollowingDirection(report []int) []int {
	increases, decreases := GetDirections(report)
	increasing := increases > decreases
	reportLen := len(report)

	for i := 0; i < reportLen; i++ {
		var res []int
		from := i
		to := i + 1
		if i == reportLen-1 {
			from = reportLen - 1
			to = reportLen - 2
			if !IsSafeNumber(from, to, !increasing) {
				return append(res, report[:to]...)
			}
		} else if !IsSafeNumber(report[from], report[to], increasing) {
			if i == reportLen-2 {
				return append(res, report[:to]...)
			}
			if IsSafeNumber(report[from], report[to+1], increasing) {
				res = append(res, report[:from+1]...)
				res = append(res, report[to+1:]...)
			} else {
				res = append(res, report[:from]...)
				res = append(res, report[to:]...)
			}
			return res
		}
	}

	return report
}

func ReportSafe_part2(report []int) bool {
	directionReport := GetReportDirectionReport(report)

	if directionReport == MoreThanOneDifferingElement || directionReport == NoDifferingElements {
		return false
	}

	report = RemoveElementNotFollowingDirection(report)
	increasing, decreasing := GetDirections(report)
	isIncreasing := increasing > decreasing

	for i := 1; i < len(report); i++ {
		if !IsSafeNumber(report[i-1], report[i], isIncreasing) {
			return false
		}
	}

	return true
}

func GetSafeReports_part2(reports [][]int) [][]int {
	var validReports [][]int

	for _, report := range reports {
		if ReportSafe_part2(report) {
			validReports = append(validReports, report)
		} else {
			fmt.Println(report)
		}
	}

	return validReports
}

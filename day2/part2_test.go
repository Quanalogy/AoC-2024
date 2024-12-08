package main

import (
	"reflect"
	"testing"
)

func Test_ReportSafe_part2_sequences_differing_by_1_to_3_should_be_true(t *testing.T) {
	reports := [][]int{
		{1, 2, 3, 4},
		{1, 3, 5, 7, 9},
		{1, 4, 7, 10, 13},
		{4, 3, 2, 1},
		{9, 7, 5, 3, 1},
		{13, 10, 7, 4, 1},
	}

	for _, report := range reports {
		isSafe := ReportSafe_part2(report)

		if !isSafe {
			t.Errorf("Expected true, got false in %v for report %v", isSafe, report)
		}
	}
}

func Test_ReportSafe_part2_sequences_differing_by_0_should_be_false(t *testing.T) {
	reports := [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{10, 10, 10, 10, 10},
	}

	for _, report := range reports {
		isSafe := ReportSafe_part2(report)
		if isSafe {
			t.Errorf("Expected false, got true in %v for report %v", isSafe, report)
		}
	}
}

type TestPair struct {
	input    []int
	expected []int
}

func Test_RemoveUnsafeElement_sequences_with_one_wrong_number_will_be_returned_without_it(t *testing.T) {
	testPairs := []TestPair{
		{[]int{1, 2, 3, 17, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 3, 10, 5, 7, 9}, []int{1, 3, 5, 7, 9}},
		{[]int{1, 4, 8, 7, 10, 13}, []int{1, 4, 7, 10, 13}},
		{[]int{4, 5, 3, 2, 1}, []int{4, 3, 2, 1}},
		{[]int{6, 7, 5, 3, 1}, []int{6, 5, 3, 1}},
		{[]int{13, 10, 7, 4, 5}, []int{13, 10, 7, 4}},
		{[]int{54, 58, 61, 64, 66}, []int{58, 61, 64, 66}},
		{[]int{48, 46, 44, 43, 42, 43}, []int{48, 46, 44, 43, 42}},
		{[]int{1, 10, 4, 7, 10, 13}, []int{1, 4, 7, 10, 13}},
	}

	for _, testPair := range testPairs {
		report := testPair.input
		actual := RemoveElementNotFollowingDirection(report)
		expected := testPair.expected

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, got %v for report %v", expected, actual, report)
		}
	}
}

func Test_GetSafeReports_part2_should_return_all_reports_that_are_safe(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	actual := GetSafeReports_part2(reports)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

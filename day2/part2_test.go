package main

import "testing"

func TestGetSafe_part2_sequences_differing_by_1_to_3_should_be_true(t *testing.T) {
	reports := [][]int{
		{1, 2, 3, 4},
		{1, 3, 5, 7, 9},
		{1, 4, 7, 10, 13},
		{4, 3, 2, 1},
		{9, 7, 5, 3, 1},
		{13, 10, 7, 4, 1},
	}

	for _, report := range reports {
		isSafe, safeCount := GetSafe_part2(report)
		if safeCount != len(report)-1 {
			t.Errorf("Expected %d, got %d", len(report), safeCount)
		}
		for _, safe := range isSafe {
			if !safe {
				t.Errorf("Expected true, got false in %v for report %v", isSafe, report)
			}
		}
	}
}

func TestGetSafe_part2_sequences_differing_by_0_should_be_false(t *testing.T) {
	reports := [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{10, 10, 10, 10, 10},
	}

	for _, report := range reports {
		isSafe, safeCount := GetSafe_part2(report)
		if safeCount != -1*(len(report)-1) {
			t.Errorf("Expected %d, got %d", len(report), safeCount)
		}
		for _, safe := range isSafe {
			if safe {
				t.Errorf("Expected false, got true in %v for report %v", isSafe, report)
			}
		}
	}
}

func TestGetSafe_part2_sequences_with_one_wrong_number_should_have_safecount_of_amount_of_elements_minus_3(t *testing.T) {
	reports := [][]int{
		{1, 2, 5, 3, 4},
		{1, 3, 10, 5, 7, 9},
		// {1, 4, 8, 7, 10, 13},
		// {4, 5, 3, 2, 1},
		// {6, 7, 5, 3, 1},
		// {13, 10, 7, 4, 5},
	}

	for _, report := range reports {
		_, safeCount := GetSafe_part2(report)
		expectedSafeCount := len(report) - 3
		if safeCount != expectedSafeCount {
			t.Errorf("Expected %d, got %d for report %v", expectedSafeCount, safeCount, report)
		}
	}
}

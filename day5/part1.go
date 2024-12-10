package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetPageOrderingRules(input string) map[int][]int {
	pageOrderingRules := make(map[int][]int)
	lines := strings.Split(input, "\n")
	numbersRegex := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		matches := numbersRegex.FindAllString(line, -1)
		page, _ := strconv.Atoi(matches[0])
		rule, _ := strconv.Atoi(matches[1])
		pageOrderingRules[page] = append(pageOrderingRules[page], rule)
	}

	return pageOrderingRules
}

func GetPageOrdering(input string) [][]int {
	pageOrdering := [][]int{}
	lines := strings.Split(input, "\n")
	numbersRegex := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		matches := numbersRegex.FindAllString(line, -1)
		page := make([]int, len(matches))
		for i, match := range matches {
			page[i], _ = strconv.Atoi(match)
		}
		pageOrdering = append(pageOrdering, page)
	}

	return pageOrdering
}

func ValidPageOrdering(pageOrdering []int, pageOrderingRules map[int][]int) bool {
	for i, page := range pageOrdering {
		rules := pageOrderingRules[page]
		if rules == nil {
			continue
		}
		if i == 0 {
			continue
		}

		for j := 0; j < i; j++ {
			prevPage := pageOrdering[j]
			if slices.Contains(rules, prevPage) {
				return false
			}
		}
	}

	return true
}

func GetValidPageOrderings(pageOrderings [][]int, pageOrderingRules map[int][]int) [][]int {
	validPageOrderings := [][]int{}

	for _, pageOrdering := range pageOrderings {
		if ValidPageOrdering(pageOrdering, pageOrderingRules) {
			validPageOrderings = append(validPageOrderings, pageOrdering)
		}
	}

	return validPageOrderings
}

func GetMiddleNumberJoinedOfValidPages(pageOrderings [][]int, pageOrderingRules map[int][]int) int {
	validPageOrderings := GetValidPageOrderings(pageOrderings, pageOrderingRules)
	res := 0

	for _, validPage := range validPageOrderings {
		res += validPage[len(validPage)/2]
	}

	return res
}

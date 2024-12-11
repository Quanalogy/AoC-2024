package main

import "slices"

func GetInvalidPageOrderings(pageOrderings [][]int, pageOrderingRules map[int][]int) [][]int {
	invalidPageOrderings := [][]int{}

	for _, pageOrdering := range pageOrderings {
		if !ValidPageOrdering(pageOrdering, pageOrderingRules) {
			invalidPageOrderings = append(invalidPageOrderings, pageOrdering)
		}
	}

	return invalidPageOrderings
}

func OrderInvalidPageOrderings(invalidPageOrderings [][]int, pageOrderingRules map[int][]int) [][]int {
	var res [][]int

	for _, invalidPageOrdering := range invalidPageOrderings {
		var pageOrdering []int

	reorder:
		for _, page := range invalidPageOrdering {
			pageRules := pageOrderingRules[page]

			for k, existingPage := range pageOrdering {
				if slices.Contains(pageRules, existingPage) {
					pageOrdering = append(pageOrdering[:k], append([]int{page}, pageOrdering[k:]...)...)
					continue reorder
				}
			}
			pageOrdering = append(pageOrdering, page)
		}
		res = append(res, pageOrdering)
	}

	return res
}

func GetMiddleNumberJoinedOfInvalidPages(pageOrderings [][]int, pageOrderingRules map[int][]int) int {
	invalidPageOrderings := GetInvalidPageOrderings(pageOrderings, pageOrderingRules)
	validPageOrderings := OrderInvalidPageOrderings(invalidPageOrderings, pageOrderingRules)
	res := 0

	for _, validPage := range validPageOrderings {
		res += validPage[len(validPage)/2]
	}

	return res
}

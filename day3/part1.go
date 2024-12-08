package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Multiply(input string) int {
	mulRegexp := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	mulMatches := mulRegexp.FindAllStringSubmatch(input, -1)

	res := 0

	for _, match := range mulMatches {
		nums := strings.Split(match[1], ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		res += num1 * num2
	}

	return res
}

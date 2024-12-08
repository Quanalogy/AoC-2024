package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Multiply_part2(input string) int {
	mulRegexp := regexp.MustCompile(`mul\((\d+,\d+)\)|do\(\)|don't\(\)`)
	mulMatches := mulRegexp.FindAllStringSubmatch(input, -1)

	res := 0
	do := true

	for _, match := range mulMatches {
		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		}

		if !do {
			continue
		}

		nums := strings.Split(match[1], ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		res += num1 * num2
	}

	return res
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetDirectionInputOutput struct {
	from     Point
	to       Point
	expected Direction
}

func Test_GetDirection(t *testing.T) {
	inputOutput := []GetDirectionInputOutput{
		{Point{0, 0}, Point{0, 1}, Down},
		{Point{0, 0}, Point{1, 0}, Right},
		{Point{1, 1}, Point{0, 1}, Left},
		{Point{1, 1}, Point{1, 0}, Up},
		{Point{1, 1}, Point{0, 0}, VirticalUpLeft},
		{Point{1, 1}, Point{2, 0}, VirticalUpRight},
		{Point{1, 1}, Point{0, 2}, VirticalDownLeft},
		{Point{1, 1}, Point{2, 2}, VirticalDownRight},
		{Point{3, 9}, Point{2, 8}, VirticalUpLeft},
		{Point{2, 8}, Point{1, 7}, VirticalUpLeft},
		{Point{1, 7}, Point{0, 6}, VirticalUpLeft},
	}

	for _, io := range inputOutput {
		actual := GetDirection(io.from, io.to)
		assert.Equal(t, io.expected, actual, "From (x,y): %v, To: %v should be: %v", io.from, io.to, io.expected)
	}
}

func Test_Xmas_part1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	expected := 18

	actual := Xmas(input)

	assert.Equal(t, expected, actual)
}

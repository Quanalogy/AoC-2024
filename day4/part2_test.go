package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Xmas_part2(t *testing.T) {
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
	expected := 9

	actual := Xmas_part2(input)

	assert.Equal(t, expected, actual)
}

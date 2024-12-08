package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multiply_part2_example_input(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := 48

	actual := Multiply_part2(input)

	assert.Equal(t, expected, actual)
}

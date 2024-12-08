package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multiply_example_input(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	expected := 161

	actual := Multiply(input)

	assert.Equal(t, expected, actual)
}

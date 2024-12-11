package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OrderInvalidPageOrderings(t *testing.T) {
	input := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	expected := [][]int{
		{97, 75, 47, 61, 53},
		{61, 29, 13},
		{97, 75, 47, 29, 13},
	}
	rules := map[int][]int{
		47: {29, 61, 13, 53},
		97: {75, 53, 29, 13, 61, 47},
		75: {13, 61, 47, 53, 29},
		61: {29, 53, 13},
		29: {13},
		53: {13, 29},
	}

	actual := OrderInvalidPageOrderings(input, rules)

	assert.EqualValues(t, expected, actual)
}

func Test_GetMiddleNumberJoinedOfInvalidPages(t *testing.T) {
	input := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	rules := map[int][]int{
		47: {29, 61, 13, 53},
		97: {75, 53, 29, 13, 61, 47},
		75: {13, 61, 47, 53, 29},
		61: {29, 53, 13},
		29: {13},
		53: {13, 29},
	}
	expected := 123

	actual := GetMiddleNumberJoinedOfInvalidPages(input, rules)

	assert.Equal(t, expected, actual)
}

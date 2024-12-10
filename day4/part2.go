package main

import (
	"errors"
	"strings"
)

func GetDiagonalCharacters(inputLines []string, point Point) ([]string, error) {
	if point.x == 0 || point.y == 0 || point.y == len(inputLines)-1 || point.x == len(inputLines[0])-1 {
		return nil, errors.New("Point is on the edge")
	}
	return []string{
		string(inputLines[point.y-1][point.x-1]),
		string(inputLines[point.y-1][point.x+1]),
		string(inputLines[point.y+1][point.x-1]),
		string(inputLines[point.y+1][point.x+1]),
	}, nil
}

func IsX_Mas(point Point, inputLines []string) bool {
	diagonalCharacters, err := GetDiagonalCharacters(inputLines, point)
	if err != nil {
		return false
	}

	if diagonalCharacters[0] == diagonalCharacters[3] || diagonalCharacters[1] == diagonalCharacters[2] {
		return false
	}

	amountM := 0
	amountS := 0

	for _, character := range diagonalCharacters {
		if character == "M" {
			amountM += 1
		}
		if character == "S" {
			amountS += 1
		}
	}
	return amountM == 2 && amountS == 2
}

// debug function for seeing that you need to make sure that diagonally it is not MM and SS.
func PrintXMas(point Point, inputLines []string) {
	println(
		inputLines[point.y][point.x],
		string(inputLines[point.y-1][point.x-1]),
		string(inputLines[point.y-1][point.x+1]),
		string(inputLines[point.y+1][point.x-1]),
		string(inputLines[point.y+1][point.x+1]),
	)
}

func Xmas_part2(input string) int {
	inputLines := strings.Split(input, "\n")
	res := 0

	for y, line := range inputLines {
		for x, character := range line {
			if character != 'A' {
				continue
			}

			if IsX_Mas(Point{x, y}, inputLines) {
				res += 1
			}
		}
	}
	return res
}

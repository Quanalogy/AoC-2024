package main

import "strings"

type Point struct {
	x int
	y int
}

func HasChacterAtPoint(inputLines []string, point Point, character string) bool {
	return string(inputLines[point.y][point.x]) == character
}

func GetPointsAroundWithCharacter(inputLines []string, point Point, character string) []Point {
	res := []Point{}
	pointsToCheck := []Point{}
	xMin := point.x-1 >= 0
	xMax := point.x+1 < len(inputLines[0])
	yMin := point.y-1 > 0
	yMax := point.y+1 < len(inputLines)

	// above
	if yMin {
		if xMin {
			pointsToCheck = append(pointsToCheck, Point{point.x - 1, point.y - 1})
		}
		pointsToCheck = append(pointsToCheck, Point{point.x, point.y - 1})
		if xMax {
			pointsToCheck = append(pointsToCheck, Point{point.x + 1, point.y - 1})
		}
	}

	// same line
	if xMin {
		pointsToCheck = append(pointsToCheck, Point{point.x - 1, point.y})
	}
	if xMax {
		pointsToCheck = append(pointsToCheck, Point{point.x + 1, point.y})
	}

	// below
	if yMax {
		if xMin {
			pointsToCheck = append(pointsToCheck, Point{point.x - 1, point.y + 1})
		}
		pointsToCheck = append(pointsToCheck, Point{point.x, point.y + 1})
		if xMax {
			pointsToCheck = append(pointsToCheck, Point{point.x + 1, point.y + 1})
		}
	}

	for _, pointToCheck := range pointsToCheck {
		if HasChacterAtPoint(inputLines, pointToCheck, character) {
			res = append(res, pointToCheck)
		}
	}

	return res
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	VirticalUpRight
	VirticalUpLeft
	VirticalDownRight
	VirticalDownLeft
)

func GetDirection(from, to Point) Direction {
	if from.y < to.y {
		if from.x < to.x {
			return VirticalDownRight
		}
		if from.x > to.x {
			return VirticalDownLeft
		}
		return Down
	} else if from.y > to.y {
		if from.x < to.x {
			return VirticalUpRight
		}
		if from.x > to.x {
			return VirticalUpLeft
		}
		return Up
	} else {
		if from.x < to.x {
			return Right
		} else if from.x > to.x {
			return Left
		}
		panic("from and to are the same point")
	}
}

func Xmas(input string) int {
	inputLines := strings.Split(input, "\n")
	res := 0

	for y, line := range inputLines {
		for x, character := range line {
			if character != 'X' {
				continue
			}

			XPoint := Point{x, y}
			MPoints := GetPointsAroundWithCharacter(inputLines, XPoint, "M")
			if len(MPoints) == 0 {
				continue
			}

			for _, MPoint := range MPoints {
				direction := GetDirection(XPoint, MPoint)
				APoints := GetPointsAroundWithCharacter(inputLines, MPoint, "A")
				if len(APoints) == 0 {
					continue
				}

				for _, APoint := range APoints {
					if GetDirection(MPoint, APoint) != direction {
						continue
					}
					SPoints := GetPointsAroundWithCharacter(inputLines, APoint, "S")
					if len(SPoints) == 0 {
						continue
					}

					for _, SPoint := range SPoints {
						if GetDirection(APoint, SPoint) != direction {
							continue
						}
						res += 1
					}
				}
			}
		}
	}
	return res
}

package main

import (
	"advent-of-code-2024/util"
	_ "embed"
	"fmt"
	"regexp"
	"slices"
)

//go:embed test4.txt
var test []byte

//go:embed input04.txt
var input []byte

func main() {
	//lines := util.ReadLines(test)
	lines := util.ReadLines(input)
	matrix := make([][]byte, len(lines))
	for i, line := range lines {
		matrix[i] = []byte(line)
	}

	solvePart1(matrix)
	solvePart2(matrix)
}

func solvePart1(matrix [][]byte) {
	invertedMatrix := make([][]byte, len(matrix))
	for y, line := range matrix {
		invertedMatrix[y] = make([]byte, len(matrix[y]))
		copy(invertedMatrix[y], line)
		slices.Reverse(invertedMatrix[y])
	}
	slices.Reverse(invertedMatrix)

	found := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			found += checkXmasPart1(matrix, x, y)
			found += checkXmasPart1(invertedMatrix, x, y)
		}
	}
	fmt.Println(found)
}

func checkXmasPart1(matrix [][]byte, x int, y int) int {
	if matrix[y][x] != 'X' {
		return 0
	}

	height, width := len(matrix), len(matrix[0])
	dy, dx := height-y, width-x

	found := 0

	if dx >= 4 && y >= 3 && matrix[y-1][x+1] == 'M' && matrix[y-2][x+2] == 'A' && matrix[y-3][x+3] == 'S' {
		found++
	}

	if dx >= 4 && dy >= 4 && matrix[y+1][x+1] == 'M' && matrix[y+2][x+2] == 'A' && matrix[y+3][x+3] == 'S' {
		found++
	}

	if dx >= 4 && matrix[y][x+1] == 'M' && matrix[y][x+2] == 'A' && matrix[y][x+3] == 'S' {
		found++
	}

	if dy >= 4 && matrix[y+1][x] == 'M' && matrix[y+2][x] == 'A' && matrix[y+3][x] == 'S' {
		found++
	}

	return found
}

func solvePart2(matrix [][]byte) {
	found := 0

	regex := regexp.MustCompile("MMSS")

	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			if matrix[y][x] != 'A' {
				continue
			}

			topLeft := matrix[y-1][x-1]
			topRight := matrix[y-1][x+1]
			bottomLeft := matrix[y+1][x-1]
			bottomRight := matrix[y+1][x+1]

			corners := string([]byte{topLeft, topRight, bottomRight, bottomLeft})
			corners += corners
			if regex.Match([]byte(corners)) {
				found++
			}
		}
	}
	fmt.Println(found)
}

package main

import (
	"advent-of-code-2024/util"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed test3.txt
var test []byte

//go:embed input03.txt
var input []byte

func main() {
	//lines := util.ReadLines(test)
	lines := util.ReadLines(input)

	solvePart1(lines)
	//solvePart2(lines)
}

func solvePart1(lines []string) {

	regex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	sum := 0

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			leftVal, _ := strconv.Atoi(match[1])
			rightVal, _ := strconv.Atoi(match[2])
			sum += leftVal * rightVal
		}
	}
	fmt.Println(sum)
}

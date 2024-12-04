package main

import (
	"advent-of-code-2024/util"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed test3.txt
var test []byte

//go:embed input03.txt
var input []byte

func main() {
	//lines := util.ReadLines(test)
	lines := util.ReadLines(input)

	solvePart1(lines)
	solvePart2(lines)
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

func solvePart2(lines []string) {
	input := "do()" + strings.Join(lines, "")

	input = strings.ReplaceAll(input, "do()", "|||||do()")
	input = strings.ReplaceAll(input, "don't()", "|||||don't()")

	lines2 := strings.Split(input, "|||||")
	lines3 := lines2[:0]

	for _, line := range lines2 {
		if strings.HasPrefix(line, "do()") {
			lines3 = append(lines3, line)
		}
	}

	solvePart1(lines3)
}

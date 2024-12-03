package main

import (
	"advent-of-code-2024/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test2.txt
var test2 []byte

//go:embed input02.txt
var input []byte

func main() {
	//lines := util.ReadLines(test2)
	lines := util.ReadLines(input)

	solvePart1(lines)
	//solvePart2(lines)
}

func solvePart1(lines []string) {

	safeLines := 0

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		levels := strings.Split(line, " ")
		var lastLevel int
		var shouldAscend bool
		safe := true

		for i, levelString := range levels {
			levelParsed, _ := strconv.ParseInt(levelString, 10, 32)
			level := int(levelParsed)

			if i > 0 {
				difference := max(lastLevel, level) - min(lastLevel, level)

				if difference == 0 || difference > 3 {
					safe = false
					break
				}

				ascending := level > lastLevel

				if i == 1 {
					shouldAscend = ascending
				}

				if ascending != shouldAscend {
					safe = false
				}
			}

			lastLevel = level
		}

		if safe {
			safeLines++
		}
	}

	fmt.Println(safeLines)
}

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
	solvePart2(lines)
}

func solvePart1(lines []string) {
	solve(lines, 0)
}

func solvePart2(lines []string) {
	solve(lines, 1)
}

func solve(reports []string, dampenerStrength int) {
	safeLines := 0

	for _, report := range reports {
		if len(report) < 1 {
			continue
		}

		levelStrings := strings.Split(report, " ")
		levels := []int{}

		for _, levelString := range levelStrings {
			level, _ := strconv.Atoi(levelString)
			levels = append(levels, level)
		}

		if isReportSafe(levels, dampenerStrength) {
			safeLines++
		}
	}

	fmt.Println(safeLines)
}

func isReportSafe(levels []int, dampenerStrength int) bool {
	var lastLevel int
	var shouldAscend bool
	safe := true
	badLevels := 0

	for i, level := range levels {
		if i > 0 {
			difference := max(lastLevel, level) - min(lastLevel, level)

			if difference == 0 || difference > 3 {
				safe = false
			} else {
				ascending := level > lastLevel

				if i == 1 {
					shouldAscend = ascending
				}

				if ascending != shouldAscend {
					safe = false
				}
			}
		}

		if !safe {
			safe = true
			badLevels++
		}

		lastLevel = level
	}

	if dampenerStrength < 1 {
		return badLevels == 0
	}

	for i := range levels {
		levelsCopy := make([]int, len(levels))
		copy(levelsCopy, levels)

		if i == len(levels)-1 {
			levelsCopy = levelsCopy[:len(levels)-1]
		} else {
			copy(levelsCopy[i:], levelsCopy[i+1:])
			levelsCopy = levelsCopy[:len(levelsCopy)-1]
		}

		if isReportSafe(levelsCopy, dampenerStrength-1) {
			return true
		}
	}

	return badLevels == 0
}

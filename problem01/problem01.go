package main

import (
	"advent-of-code-2024/util"
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

//go:embed test1.txt
var test1 []byte

//go:embed input01.txt
var input []byte

func main() {
	//lines := util.ReadLines(test1)
	lines := util.ReadLines(input)

	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) {
	left := []int{}
	right := []int{}

	re := regexp.MustCompile(" +")

	for _, line := range lines {
		if len(line) > 0 {
			split := re.Split(line, -1)
			leftValue, _ := strconv.ParseInt(split[0], 10, 32)
			rightValue, _ := strconv.ParseInt(split[1], 10, 32)

			left = append(left, int(leftValue))
			right = append(right, int(rightValue))
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += max(right[i], left[i]) - min(right[i], left[i])
	}

	fmt.Println(sum)
}

func solvePart2(lines []string) {
	leftCountsInRight := make(map[int]int)
	leftCountsInLeft := make(map[int]int)

	right := []int{}

	re := regexp.MustCompile(" +")

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		split := re.Split(line, -1)
		leftValue, _ := strconv.ParseInt(split[0], 10, 32)
		rightValue, _ := strconv.ParseInt(split[1], 10, 32)

		leftValueInt := int(leftValue)

		leftCountsInRight[leftValueInt] = 0

		if _, ok := leftCountsInLeft[leftValueInt]; ok {
			leftCountsInLeft[leftValueInt]++
		} else {
			leftCountsInLeft[leftValueInt] = 1
		}

		right = append(right, int(rightValue))
	}

	for _, rightValue := range right {
		if _, ok := leftCountsInRight[rightValue]; ok {
			leftCountsInRight[rightValue]++
		}
	}

	sum := 0
	for leftKey, leftValue := range leftCountsInRight {
		sum += leftKey * leftValue * leftCountsInLeft[leftKey]
	}

	fmt.Println(sum)
}

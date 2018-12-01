package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
)

func main() {
	input, err := shared.LoadInputFile("day1/input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")
	numbers := parseLines(lines)
	part1(numbers)
	part2(numbers)
}

func parseLines(lines []string) []int64 {
	numbers := make([]int64, len(lines))
	for i, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic("Input line is not a number")
		}
		numbers[i] = num
	}
	return numbers
}

func part1(numbers []int64) {
	var currTotal int64

	for _, num := range numbers {
		currTotal += num
	}

	fmt.Printf("Resulting frequency: %d\n", currTotal)
}

func part2(numbers []int64) {
	var currTotal int64
	var totals = make([]int64, 0)
	var firstDouble int64

	index := 0

	for {
		currNum := numbers[index]
		currTotal += currNum

		if contains(totals, currTotal) {
			firstDouble = currTotal
			break
		}

		totals = append(totals, currTotal)

		index++

		if index >= len(numbers) {
			index = 0
		}
	}

	fmt.Printf("Resulting frequency: %d\n", firstDouble)
}

func contains(haystack []int64, needle int64) bool {
	for _, thing := range haystack {
		if thing == needle {
			return true
		}
	}
	return false
}

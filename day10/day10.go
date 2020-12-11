package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	start := time.Now()
	bags := strings.Split(input, "\n")
	bagIndex := parseRatings(bags)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(bagIndex)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(bagIndex)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(numbers []int) shared.Result {
	ratings := make([]int, len(numbers))
	copy(ratings, numbers)
	sort.Ints(ratings)

	currJolts := 0

	diffs := make(map[int]int)

	for _, rating := range ratings {
		diff := rating - currJolts
		if diff > 0 && diff <= 3 {
			diffs[diff]++
			currJolts = rating
		} else {
			break
		}
	}

	// Final adapter
	diffs[3]++

	return func() {
		fmt.Printf("%d * %d = %d\n", diffs[1], diffs[3], diffs[1]*diffs[3])
	}
}

func part2(numbers []int) shared.Result {
	return shared.NoopResult
}

func parseRatings(lines []string) []int {
	ratings := make([]int, len(lines))

	for i, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		ratings[i] = num
	}

	return ratings
}

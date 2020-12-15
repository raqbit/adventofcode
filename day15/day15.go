package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
	"time"
)

const (
	Test  = "3,1,2"
	Input = "18,11,9,0,5,1"
)

func main() {
	start := time.Now()
	nums := parseNums(Input)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(nums)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(nums)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(startingNums []int) shared.Result {
	result := game(startingNums, 2020)

	return func() {
		fmt.Printf("Last spoken: %d\n", result)
	}
}

func part2(nums []int) shared.Result {
	result := game(nums, 30000000)
	return func() {
		fmt.Printf("Last spoken: %d\n", result)
	}
}

func game(startingNums []int, end int) int {
	spokenNums := make(map[int]int)
	lastSpoken := -1

	for turn := 0; turn < end; turn++ {
		num := 0

		if turn < len(startingNums) {
			num = startingNums[turn]
			spokenNums[num] = turn
			lastSpoken = num
			continue
		}

		if whenLastSpoken, ok := spokenNums[lastSpoken]; ok {
			num = turn - 1 - whenLastSpoken
		}

		spokenNums[lastSpoken] = turn - 1

		lastSpoken = num
	}

	return lastSpoken
}

func parseNums(input string) []int {
	parts := strings.Split(input, ",")
	nums := make([]int, len(parts))

	for i, part := range parts {
		nums[i], _ = strconv.Atoi(part)
	}

	return nums
}

package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
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
	bagIndex := parseNumbers(bags)
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
	num := findInvalidNum(numbers, 25)

	return func() {
		fmt.Printf("Invalid number: %d\n", num)
	}
}

func findInvalidNum(numbers []int, preambleLen int) int {
nums:
	for i := preambleLen + 1; i < len(numbers); i++ {
		num := numbers[i]

		// Naive approach
		// TODO: improve
		for j := i - preambleLen; j < i; j++ {
			for k := i - preambleLen; k < i; k++ {
				if j == k {
					continue
				}

				a := numbers[j]
				b := numbers[k]

				if a == b {
					continue
				}

				if num == a+b {
					// Valid
					continue nums
				}
			}
		}

		// Invalid
		return num
	}

	return -1
}

func part2(numbers []int) shared.Result {
	num := findInvalidNum(numbers, 25)

	i := 0
	j := 1

	smallest := math.MaxInt64
	largest := 0

	for {
		if j >= len(numbers) {
			return shared.NoopResult
		}

		total := 0
		smallest = math.MaxInt64
		largest = 0

		for k := i; k <= j; k++ {
			if numbers[k] < smallest {
				smallest = numbers[k]
			}

			if numbers[k] > largest {
				largest = numbers[k]
			}

			total += numbers[k]
		}

		if total == num {
			break
		}

		if total > num {
			i++
		} else {
			j++
		}
	}

	return func() {
		fmt.Printf("Weakness: %d\n", smallest+largest)
	}
}

func parseNumbers(lines []string) []int {
	numbers := make([]int, len(lines))

	for i, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		numbers[i] = num
	}

	return numbers
}

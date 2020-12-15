package main

import (
	"fmt"
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
	earliestTime, busses := parseNotes(bags)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(earliestTime, busses)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(busses)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(earliestTime int, busses []int) shared.Result {
	for currentTime := earliestTime; true; currentTime++ {
		for _, id := range busses {
			if id == 0 {
				continue
			}
			if currentTime%id == 0 {
				return func() {
					fmt.Printf("%d*(%d-%d) = %d\n", id, currentTime, earliestTime, id*(currentTime-earliestTime))
				}
			}
		}
	}

	return shared.NoopResult
}

func part2(busses []int) shared.Result {
	highestIntervalOffset := 0
	highestInterval := 0
	startTime := 100000000000000
	//startTime := 0

	offsetMap := make(map[int]int)

	for i, bus := range busses {
		if bus != 0 {
			offsetMap[i] = bus
		}

		if bus > highestInterval {
			highestInterval = bus
			highestIntervalOffset = i
		}
	}

	iterations := 0

	// Go in increments of the bus with highest interval
	// FIXME: Brute force isn't the answer, sadly
	for timestamp := startTime; true; timestamp += highestInterval {
		firstBus := timestamp - highestIntervalOffset
		valid := true

		for offset, bus := range offsetMap {
			if bus == highestInterval {
				continue
			}

			if (firstBus+offset)%bus != 0 {
				valid = false
				break
			}
		}

		if valid {
			return func() {
				fmt.Printf("First timestamp: %d, iterations: %d * 1889 (%d)\n", firstBus, iterations, timestamp)
			}
		}

		iterations++
	}

	return shared.NoopResult
}

func parseNotes(lines []string) (int, []int) {
	earliestTime, _ := strconv.Atoi(lines[0])

	nums := strings.Split(lines[1], ",")
	busses := make([]int, len(nums))

	for i, num := range nums {
		if num == "x" {
			busses[i] = 0
		}

		value, _ := strconv.Atoi(num)

		busses[i] = value
	}

	return earliestTime, busses
}

package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")
	numbers := parseLines(lines)
	part1(numbers)
	fmt.Println("----")
	part2(numbers)
}

func part1(moduleMasses []int) {
	var fuelTotal int

	// Add calculated fuel for all modules to the total
	for _, moduleMass := range moduleMasses {
		fuelTotal += calculateFuel(moduleMass)
	}

	fmt.Printf("Rocket total fuel mass: %d\n", fuelTotal)
}

func part2(moduleMasses []int) {
	var rocketTotal int

	for _, moduleMass := range moduleMasses {
		var moduleTotal int
		remainingMass := moduleMass

		for {
			fuelNeed := calculateFuel(remainingMass)
			if fuelNeed <= 0 {
				// Wish really hard
				break
			}
			moduleTotal += fuelNeed
			remainingMass = fuelNeed
		}

		// Add module total to rocket total
		rocketTotal += moduleTotal
	}

	fmt.Printf("Rocket total fuel mass: %d\n", rocketTotal)
}

// Calculates the fuel need for a given mass
func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass/3.0))) - 2
}

// Parses the input lines as integers
func parseLines(lines []string) []int {
	numbers := make([]int, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic("Input line is not a number")
		}
		numbers[i] = int(num)
	}
	return numbers
}

package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strings"
	"time"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")
	landMap := parseLines(lines)

	start := time.Now()
	res := part1(landMap)
	duration := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", duration.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(landMap)
	duration = time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", duration.String())
}

func part1(landMap [][]bool) shared.Result {
	// Get tree count for 3,1 slope
	treeCount := checkSlope(landMap, 3, 1)

	return func() {
		fmt.Printf("Hit %d trees\n", treeCount)
	}
}

func part2(landMap [][]bool) shared.Result {
	// All given slopes
	slopes := []struct {
		dX int
		dY int
	}{
		{
			dX: 1,
			dY: 1,
		},
		{
			dX: 3,
			dY: 1,
		},
		{
			dX: 5,
			dY: 1,
		},
		{
			dX: 7,
			dY: 1,
		},
		{
			dX: 1,
			dY: 2,
		},
	}

	product := 1

	for _, slope := range slopes {
		// Calculate product of the trees found for all slopes
		product *= checkSlope(landMap, slope.dX, slope.dY)
	}

	return func() {
		fmt.Printf("Product: %d\n", product)
	}
}

func checkSlope(landMap [][]bool, dX int, dY int) int {
	height := len(landMap)
	width := len(landMap[0])

	// Map coordinates
	x := 0
	y := 0

	// Amount of trees hit
	treeCount := 0

	for {
		// Make move
		x += dX
		y += dY

		// Stop counting if we reached the bottom
		if y >= height {
			break
		}

		// Wrap around map
		x = x % width

		// Has tree
		if landMap[y][x] {
			treeCount++
		}
	}

	return treeCount
}

func parseLines(lines []string) [][]bool {
	landMap := make([][]bool, len(lines)-1)

	for i, line := range lines {
		if line == "" {
			continue
		}

		row := make([]bool, len(line))
		for j, char := range line {
			switch char {
			case '#':
				row[j] = true
			default:
				fallthrough
			case '.':
				row[j] = false
			}
		}
		landMap[i] = row
	}

	return landMap
}

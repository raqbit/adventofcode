package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"regexp"
	"strconv"
	"strings"
)

type square struct {
	x  int64
	y  int64
	dx int64
	dy int64
}

const (
	fabricSize = 1000
	claimRegex = `#(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<dx>\d+)x(?P<dy>\d+)`
)

func main() {
	input, err := shared.LoadInputFile("day3/input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	squares := parseClaims(lines)

	part1(squares)
	part2(squares)
}

func parseClaims(lines []string) []square {
	claimMatcher, err := regexp.Compile(claimRegex)

	if err != nil {
		panic("Could not compile claim regex")
	}

	squares := make([]square, len(lines))

	for i, line := range lines {

		parsed := parseClaim(claimMatcher, line)

		// Ignoring errors here cause I don't wanna handle all of them >.>
		x, _ := strconv.ParseInt(parsed["x"], 10, 64)
		y, _ := strconv.ParseInt(parsed["y"], 10, 64)
		dx, _ := strconv.ParseInt(parsed["dx"], 10, 64)
		dy, _ := strconv.ParseInt(parsed["dy"], 10, 64)

		squares[i] = square{
			x:  x,
			y:  y,
			dx: dx,
			dy: dy,
		}
	}

	return squares
}

func makeGrid() [][]int64 {
	grid := make([][]int64, fabricSize)

	for i := range grid {
		grid[i] = make([]int64, fabricSize)
	}

	return grid
}

func part1(squares []square) {
	grid := makeGrid()

	inchesWithMultiple := 0

	for _, sq := range squares {
		for i := sq.x; i < sq.x+sq.dx; i++ {
			for j := sq.y; j < sq.y+sq.dy; j++ {
				grid[i][j]++

				// If it's 2 we up the counter, we don't care if it's >2 since it has already been counted
				if grid[i][j] == 2 {
					inchesWithMultiple++
				}
			}
		}
	}

	fmt.Printf("Number of Inches with multiple claims: %d\n", inchesWithMultiple)
}

func part2(squares []square) {
	grid := makeGrid()

	for _, sq := range squares {
		for i := sq.x; i < sq.x+sq.dx; i++ {
			for j := sq.y; j < sq.y+sq.dy; j++ {
				grid[i][j]++
			}
		}
	}
	for i, square := range squares {
		if !squareHasOverlap(grid, square) {
			fmt.Printf("Found square without overlap: id %d\n", i+1)
		}
	}
}

func squareHasOverlap(grid [][]int64, sq square) bool {
	expectedTotal := sq.dx * sq.dy
	var total int64
	for i := sq.x; i < sq.x+sq.dx; i++ {
		for j := sq.y; j < sq.y+sq.dy; j++ {
			total += grid[i][j]
		}
	}

	return total != expectedTotal
}

func parseClaim(claimMatcher *regexp.Regexp, line string) map[string]string {
	match := claimMatcher.FindStringSubmatch(line)
	result := make(map[string]string)
	for i, name := range claimMatcher.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}

//func visualizeGrid(grid [][]int64) {
//	for _, row := range grid {
//		for _, column := range row {
//			fmt.Print(column)
//		}
//		fmt.Print("\n")
//	}
//}

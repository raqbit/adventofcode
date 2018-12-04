package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"regexp"
	"strconv"
	"strings"
)

type claim struct {
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

	claims := parseClaims(lines)

	part1(claims)
	part2(claims)
}

func parseClaims(lines []string) []claim {
	claimMatcher, err := regexp.Compile(claimRegex)

	if err != nil {
		panic("Could not compile claim regex")
	}

	claims := make([]claim, len(lines))

	for i, line := range lines {

		parsed := parseClaim(claimMatcher, line)

		// Ignoring errors here cause I don't wanna handle all of them >.>
		x, _ := strconv.ParseInt(parsed["x"], 10, 64)
		y, _ := strconv.ParseInt(parsed["y"], 10, 64)
		dx, _ := strconv.ParseInt(parsed["dx"], 10, 64)
		dy, _ := strconv.ParseInt(parsed["dy"], 10, 64)

		claims[i] = claim{
			x:  x,
			y:  y,
			dx: dx,
			dy: dy,
		}
	}

	return claims
}

func makeGrid() [][]int64 {
	grid := make([][]int64, fabricSize)

	for i := range grid {
		grid[i] = make([]int64, fabricSize)
	}

	return grid
}

func part1(claims []claim) {
	grid := makeGrid()

	inchesWithMultiple := 0

	for _, cl := range claims {
		for i := cl.x; i < cl.x+cl.dx; i++ {
			for j := cl.y; j < cl.y+cl.dy; j++ {
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

func part2(claims []claim) {
	grid := makeGrid()

	for _, cl := range claims {
		for i := cl.x; i < cl.x+cl.dx; i++ {
			for j := cl.y; j < cl.y+cl.dy; j++ {
				grid[i][j]++
			}
		}
	}
	for i, cl := range claims {
		if !claimHasOverlap(grid, cl) {
			fmt.Printf("Found claim without overlap: #%d\n", i+1)
		}
	}
}

func claimHasOverlap(grid [][]int64, cl claim) bool {
	expectedTotal := cl.dx * cl.dy
	var total int64
	for i := cl.x; i < cl.x+cl.dx; i++ {
		for j := cl.y; j < cl.y+cl.dy; j++ {
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

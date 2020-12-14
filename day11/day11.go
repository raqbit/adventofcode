package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strings"
	"time"
)

type cellType rune

const (
	Floor    cellType = '.'
	Empty    cellType = 'L'
	Occupied cellType = '#'
)

type countFunc func(grid [][]cellType, x, y int) int

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	start := time.Now()
	lines := strings.Split(input, "\n")
	grid := parseSeats(lines)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(grid)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(grid)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(grid [][]cellType) shared.Result {
	occupied := runAutomata(grid, 4, countAdjacentOccupiedChairs)

	return func() {
		fmt.Printf("# of chairs occupied: %d\n", occupied)
	}
}

func part2(grid [][]cellType) shared.Result {
	occupied := runAutomata(grid, 5, traceOccupiedChairs)

	return func() {
		fmt.Printf("# of chairs occupied: %d\n", occupied)
	}
}

func runAutomata(grid [][]cellType, deathCount int, cFunc countFunc) int {
	occupied := 0

	for {
		// Copy current grid as changes should all happen in single frame of time
		nextGrid := make([][]cellType, len(grid))
		for v, row := range grid {
			nextGrid[v] = make([]cellType, len(row))
			copy(nextGrid[v], row)
		}

		diff := false

		for y := range grid {
			for x := range grid[y] {
				switch grid[y][x] {
				case Empty:
					if (cFunc(grid, x, y)) == 0 {
						nextGrid[y][x] = Occupied
						diff = true
						occupied++
					}
				case Occupied:
					if (cFunc(grid, x, y)) >= deathCount {
						nextGrid[y][x] = Empty
						diff = true
						occupied--
					}
				}
			}
		}

		if !diff {
			return occupied
		}

		// Override grid with new one
		grid = nextGrid
	}
}

func countAdjacentOccupiedChairs(grid [][]cellType, x, y int) (count int) {
	for i := y - 1; i <= y+1; i++ {
		if i < 0 || i >= len(grid) {
			continue
		}
		for j := x - 1; j <= x+1; j++ {
			if j < 0 || j >= len(grid[i]) || (i == y && j == x) {
				continue
			}

			if grid[i][j] == Occupied {
				count++
			}
		}
	}
	return
}

func traceOccupiedChairs(grid [][]cellType, x, y int) (count int) {
	for dY := -1; dY < 2; dY++ {
		for dX := -1; dX < 2; dX++ {
			if dY == 0 && dX == 0 {
				continue
			}

			for rayLen := 1; true; rayLen++ {
				// Calculate new coords based on dX, dY direction & ray length
				newX := x + dX*rayLen
				newY := y + dY*rayLen

				// Bounds check
				if (newY < 0 || newY >= len(grid)) || (newX < 0 || newX >= len(grid[0])) {
					break
				}

				cell := grid[newY][newX]

				if cell == Occupied {
					count++
					break
				} else if cell == Empty {
					break
				}
			}
		}
	}
	return
}

func parseSeats(lines []string) [][]cellType {
	grid := make([][]cellType, len(lines))

	for i, line := range lines {
		row := make([]cellType, len(line))
		for j, cell := range line {
			row[j] = cellType(cell)
		}
		grid[i] = row
	}

	return grid
}

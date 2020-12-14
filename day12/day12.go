package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
	"time"
)

type action rune

const (
	North       action = 'N'
	South       action = 'S'
	East        action = 'E'
	West        action = 'W'
	LeftRotate  action = 'L'
	RightRotate action = 'R'
	Forward     action = 'F'
)

type instruction struct {
	ac  action
	val int
}

func (i instruction) String() string {
	return fmt.Sprintf("%c%d", i.ac, i.val)
}

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	start := time.Now()
	bags := strings.Split(input, "\n")
	bagIndex := parseInstructions(bags)
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

func part1(instructions []instruction) shared.Result {
	east := 0
	north := 0
	rot := 90

	for _, i := range instructions {
		switch i.ac {
		case North:
			north += i.val
		case South:
			north -= i.val
		case East:
			east += i.val
		case West:
			east -= i.val
		case RightRotate:
			rot = (360 + rot + i.val) % 360
		case LeftRotate:
			rot = (360 + rot - i.val) % 360
		case Forward:
			switch rot {
			case 0:
				north += i.val
			case 90:
				east += i.val
			case 180:
				north -= i.val
			case 270:
				east -= i.val
			}
		}
	}

	return func() {
		fmt.Printf("East: %d, North: %d, Distance: %d\n", east, north, dist(east, north))
	}
}

func part2(instructions []instruction) shared.Result {
	wpNorth := 1
	wpEast := 10

	shipNorth := 0
	shipEast := 0

	for _, i := range instructions {
		switch i.ac {
		case North:
			wpNorth += i.val
		case South:
			wpNorth -= i.val
		case East:
			wpEast += i.val
		case West:
			wpEast -= i.val
		case RightRotate:
			wpEast, wpNorth = rotateMatrixPosition(wpEast, wpNorth, -i.val)
		case LeftRotate:
			wpEast, wpNorth = rotateMatrixPosition(wpEast, wpNorth, i.val)
		case Forward:
			shipNorth += wpNorth * i.val
			shipEast += wpEast * i.val
		}
	}

	return func() {
		fmt.Printf("Distance: %d\n", dist(shipNorth, shipEast))
	}
}

// https://en.wikipedia.org/wiki/Rotation_matrix#In_two_dimensions
func rotateMatrixPosition(x, y, degrees int) (int, int) {
	radians := float64(degrees) * (math.Pi / 180)
	cosTheta := math.Cos(radians)
	sinTheta := math.Sin(radians)
	return int(math.Round((float64(x) * cosTheta) - (float64(y) * sinTheta))),
		int(math.Round((float64(x) * sinTheta) + (float64(y) * cosTheta)))
}

func dist(a, b int) int {
	return int(math.Abs(float64(a)) + math.Abs(float64(b)))
}

func parseInstructions(lines []string) []instruction {
	grid := make([]instruction, len(lines))

	for i, line := range lines {
		value, _ := strconv.Atoi(line[1:])

		grid[i] = instruction{
			ac:  action(line[0]),
			val: value,
		}
	}

	return grid
}

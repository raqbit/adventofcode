package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// Taxicab distance
func (p point) DistTo(other point) int {
	dist := math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y))
	// Can safely cast since result of taxicab distance should never be a non-whole number
	return int(dist)
}

type distanceTo struct {
	dist       int
	pointIndex int
}

func main() {
	input, err := shared.LoadInputFile("day6/input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	points := parseCoords(lines)
	minX, minY, maxX, maxY := calculateGridSize(points)
	part1(points, minX, minY, maxX, maxY)
	fmt.Println("-----")
	part2(points, minX, minY, maxX, maxY)
}

func part1(points []point, minX, minY, maxX, maxY int) {
	// Grab width & height
	width := minX - maxX
	height := minY - maxY

	// Calculating diagonal
	// A^2 + B^2 = C^2
	// C = sqrt(A^2+B^2)
	diag := math.Sqrt(
		math.Pow(float64(width), 2) +
			math.Pow(float64(height), 2),
	)

	areas := make(map[int]int)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			here := point{x: x, y: y}

			// Keep track of all measured distances
			distances := make([]distanceTo, 0)

			// Max distance inside a rectangle is the diagonal
			minDist := int(diag)

			for k, p := range points {
				// Calculate distance to point
				dist := here.DistTo(p)

				// If this distance is less than or equal to the distance to a previous point,
				// store it as an option
				if dist <= minDist {
					minDist = dist
					distances = append(distances, distanceTo{dist: dist, pointIndex: k})
				}
			}

			// Sort distances
			sort.Slice(distances, func(i, j int) bool {
				return distances[i].dist < distances[j].dist
			})

			// Closest two points have equal distance,
			// so we don't count this point
			if len(distances) > 1 && distances[0].dist == distances[1].dist {
				// Go to next cell
				continue
			}

			// Add to the area count of the closest point
			areas[distances[0].pointIndex]++
		}
	}

	var largestArea int

	// Get largest area
	for _, area := range areas {
		if area > largestArea {
			largestArea = area
		}
	}

	fmt.Printf("Largest area: %d\n", largestArea)
}

func part2(points []point, minX, minY, maxX, maxY int) {
	var areaSize int

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			here := point{x: x, y: y}
			distTot := 0
			for _, p := range points {
				// Calculate distance to point
				distTot += here.DistTo(p)
			}

			if distTot < 10000 {
				areaSize++
			}
		}
	}

	fmt.Printf("Area size: %d", areaSize)
}

func calculateGridSize(points []point) (minX int, minY int, maxX int, maxY int) {
	xArray := make([]int, len(points))
	yArray := make([]int, len(points))

	// Split points into arrays for both axis
	for i, point := range points {
		xArray[i] = point.x
		yArray[i] = point.y
	}

	// Sort the arrays
	sort.Ints(xArray)
	sort.Ints(yArray)

	minX = xArray[0]
	minY = yArray[0]
	maxX = xArray[len(xArray)-1]
	maxY = yArray[len(yArray)-1]

	return
}

func parseCoords(lines []string) []point {
	points := make([]point, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		coords := strings.Split(line, ", ")
		xPos, _ := strconv.ParseInt(coords[0], 10, strconv.IntSize)
		yPos, _ := strconv.ParseInt(coords[1], 10, strconv.IntSize)
		points[i] = point{
			x: int(xPos),
			y: int(yPos),
		}
	}
	return points
}

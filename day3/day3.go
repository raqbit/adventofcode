package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

// Stringer
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Taxicab distance
func (p Point) DistTo(other Point) int {
	dist := math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y))
	// Can safely cast since result of taxicab distance should never be a non-whole number
	return int(dist)
}

func (p *Point) Update(dir Direction, amount int) {
	switch dir {
	case Up:
		p.Y++
	case Down:
		p.Y--
	case Left:
		p.X--
	case Right:
		p.X++
	}
}

type Direction rune

const Up Direction = 'U'
const Down Direction = 'D'
const Left Direction = 'L'
const Right Direction = 'R'

type Wire struct {
	index        int
	instructions []Instruction
}

type Instruction struct {
	direction Direction
	amount    int
}

var centerPoint = Point{
	X: 0,
	Y: 0,
}

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	wires := parseInstructions(lines)

	part1(wires)
	fmt.Println("----")
	part2(wires)
}

func part1(wires []Wire) {
	wireMap := make(map[Point]int)

	minDist := math.MaxInt32
	for _, wire := range wires {
		// Start at centerpoint
		position := centerPoint

		for _, instruction := range wire.instructions {
			// Modify position 'amount' times
			for i := 0; i < instruction.amount; i++ {

				// Modify position based on direction
				position.Update(instruction.direction, 1)

				if other, ok := wireMap[position]; ok {
					if wire.index != other {
						// Intersection!

						// Calculate distance to central point
						dist := position.DistTo(centerPoint)

						if dist < minDist {
							minDist = dist
						}
					}
				}

				// Set lineIndex in wireMap
				wireMap[position] = wire.index
			}
		}
	}

	fmt.Println(minDist)
}

func part2(wires []Wire) {
	stepCountMap := make(map[Point]int)
	globalWireMap := make(map[Point]int)

	for _, wire := range wires {
		position := centerPoint
		wireMap := make(map[Point]bool)

		stepCount := 0
		for _, instruction := range wire.instructions {
			// Modify position 'amount' times

			for i := 0; i < instruction.amount; i++ {
				// Modify position based on direction
				position.Update(instruction.direction, 1)

				// Update step count
				stepCount++

				// We haven't been here yet
				if _, ok := wireMap[position]; !ok {
					// Update step count map
					stepCountMap[position] += stepCount

					// Mark that we have been here
					wireMap[position] = true
					globalWireMap[position] += 1
				}
			}
		}
	}

	minSteps := math.MaxInt32

	for pos, count := range globalWireMap {
		if count <= 1 {
			continue
		}

		steps := stepCountMap[pos]

		if steps < minSteps {
			minSteps = steps
		}
	}

	fmt.Println(minSteps)
}

func parseInstructions(lines []string) []Wire {
	wires := make([]Wire, 0)
	for i, line := range lines {
		if line == "" {
			continue
		}

		wireInstructions := strings.Split(line, ",")
		instructions := make([]Instruction, len(wireInstructions))

		for j, inst := range wireInstructions {
			dir := inst[0]
			num, _ := strconv.ParseInt(inst[1:], 10, 32)
			instructions[j] = Instruction{
				amount:    int(num),
				direction: Direction(dir),
			}
		}

		wires = append(wires, Wire{
			index:        i,
			instructions: instructions,
		})
	}

	return wires
}

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

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	part1(lines)
	fmt.Println("----")
	//part2(lines)
}

func part1(lines []string) {
	start := time.Now()
	coordMap := make(map[shared.Point]int)

	centralPoint := shared.Point{
		X: 0,
		Y: 0,
	}

	minDist := math.MaxInt32
	for lineIndex, v := range lines {

		if v == "" {
			break
		}

		instructions := strings.Split(v, ",")
		currentCoord := centralPoint

		for _, instruction := range instructions {
			dir, num := parseInstruction(instruction)

			for i := num; i > 0; i-- {
				switch dir {
				case 'U':
					currentCoord.Y++
				case 'D':
					currentCoord.Y--
				case 'L':
					currentCoord.X--
				case 'R':
					currentCoord.X++
				}

				if other, ok := coordMap[currentCoord]; ok {
					if lineIndex != other {
						// Calculate distance to central point
						dist := currentCoord.DistTo(centralPoint)
						if dist < minDist {
							minDist = dist
						}
					}
				}

				// Set lineIndex in coordmap
				coordMap[currentCoord] = lineIndex
			}
		}
	}

	fmt.Println(minDist)
	fmt.Printf("%v\n", time.Since(start))
}

func parseInstruction(inst string) (rune, int) {
	dir := inst[0]
	rawNum, _ := strconv.ParseInt(inst[1:], 10, 32)
	num := int(rawNum)

	return rune(dir), num
}

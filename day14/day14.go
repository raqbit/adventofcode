package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	isMask bool
	addr   uint64
	mask   uint64
	value  uint64
}

var (
	memInstructionRegex = regexp.MustCompile("mem\\[(?P<addr>\\d+)\\] = (?P<val>\\d+)")
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	start := time.Now()
	lines := strings.Split(input, "\n")
	instructions := parseInstructions(lines)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(instructions)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(instructions)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(instructions []instruction) shared.Result {
	memory := make(map[uint64]uint64)

	// Mask to apply to instruction value
	var valueMask uint64 = 0xFFFFFFFFF // 36 bits

	// Bits to override in instruction value
	var override uint64 = 0x00 // 36 bits

	for _, inst := range instructions {
		if inst.isMask {
			// Set mask & override
			valueMask = inst.mask
			override = inst.value
			continue
		} else {
			// Calculate memory value based on value mask & value override bits
			memory[inst.addr] = (inst.value & valueMask) | override
		}
	}

	// Get sum of memory values
	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return func() {
		fmt.Printf("Total: %v\n", sum)
	}
}

func part2(instructions []instruction) shared.Result {
	memory := make(map[uint64]uint64)

	// Mask to apply to instruction address
	var floatingMask uint64 = 0xFFFFFFFFF // 36 bits

	// All options of masks
	var maskOpts []uint64

	// Bits to override in instruction address
	var addressOverride uint64 = 0x00 // 36 bits

	for _, inst := range instructions {
		if inst.isMask {
			// Set mask & override
			floatingMask = inst.mask

			// Calculate all mask options
			maskOpts = getMaskOptions(floatingMask)
			addressOverride = inst.value
			continue
		} else {
			addr := inst.addr | addressOverride
			// Store in memory at all addresses given by mask options
			for _, bits := range maskOpts {
				memory[addr&^floatingMask|bits] = inst.value
			}
		}
	}

	// Get sum of memory values
	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return func() {
		fmt.Printf("Total: %v\n", sum)
	}
}

func getMaskOptions(mask uint64) []uint64 {
	bitsSet := getBitsSet(mask)

	masks := make([]uint64, int(math.Pow(2, float64(bitsSet))))

	var x uint64
	var i int

	masks[i] = x
	x = ((x | ^mask) + 1) & mask

	for x != 0 {
		masks[i] = x
		x = ((x | ^mask) + 1) & mask
		i++
	}

	return masks
}

func getBitsSet(n uint64) (bitsSet int) {
	for n != 0 {
		n &= n - 1
		bitsSet++
	}
	return
}

func parseInstructions(lines []string) []instruction {
	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " = ")
		loc := parts[0]
		val := parts[1]

		if loc == "mask" {
			var mask uint64 = 0
			var value uint64 = 0

			for j, c := range val {
				if c == 'X' {
					mask |= 0b1
				} else if c == '0' {
					value |= 0b0
				} else if c == '1' {
					value |= 0b1
				}

				if j != len(val)-1 {
					mask <<= 1
					value <<= 1
				}
			}

			instructions[i] = instruction{
				isMask: true,
				mask:   mask,
				value:  value,
			}
			continue
		}

		matches := shared.RegexMatch(memInstructionRegex, line)

		addr, _ := strconv.ParseInt(matches["addr"], 10, 64)
		value, _ := strconv.ParseInt(matches["val"], 10, 64)

		// Is mem assignment
		instructions[i] = instruction{
			isMask: false,
			addr:   uint64(addr),
			value:  uint64(value),
		}
	}

	return instructions
}

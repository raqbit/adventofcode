package main

import (
	"errors"
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
	"time"
)

const (
	Cm string = "cm"
	In string = "in"
)

var (
	eyeColors = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
)

type passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

type FieldType string

const (
	Byr FieldType = "byr"
	Iyr FieldType = "iyr"
	Eyr FieldType = "eyr"
	Hgt FieldType = "hgt"
	Hcl FieldType = "hcl"
	Ecl FieldType = "ecl"
	Pid FieldType = "pid"
	Cid FieldType = "cid"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	start := time.Now()
	passports := parseLines(lines)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(passports)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(passports)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(passports []passport) shared.Result {
	valid := 0

	for _, p := range passports {
		if passportHasRequiredFields(p) {
			valid++
		}
	}

	return func() {
		fmt.Printf("Valid passports : %d\n", valid)
	}
}

func passportHasRequiredFields(pp passport) bool {
	return pp.byr != 0 &&
		pp.iyr != 0 &&
		pp.eyr != 0 &&
		pp.hgt != "" &&
		pp.hcl != "" &&
		pp.ecl != "" &&
		pp.pid != ""
}

func part2(passports []passport) shared.Result {
	valid := 0

	for _, p := range passports {
		if err := isValidPassport(p); err == nil {
			valid++
		}
	}

	return func() {
		fmt.Printf("Valid passports : %d\n", valid)
	}
}

func isValidPassport(pp passport) error {
	if !passportHasRequiredFields(pp) {
		return errors.New("missing required fields")
	}

	if pp.byr < 1920 || pp.byr > 2002 {
		return errors.New("invalid byr")
	}

	if pp.iyr < 2010 || pp.iyr > 2020 {
		return errors.New("invalid iyr")
	}

	if pp.eyr < 2020 || pp.eyr > 2030 {
		return errors.New("invalid eyr")
	}

	if err := validateHeight(pp.hgt); err != nil {
		return err
	}

	if err := validateHairColor(pp.hcl); err != nil {
		return err
	}

	if _, ok := eyeColors[pp.ecl]; !ok {
		return errors.New("invalid ecl")
	}

	if len(pp.pid) != 9 {
		return errors.New("pid too short")
	}

	return nil
}

func validateHeight(length string) error {
	// Length of 4 is minimum for a valid height value
	if len(length) < 4 {
		return errors.New("height not enough chars")
	}

	unit := length[len(length)-2:]

	val, err := parseInt(length[:len(length)-2])

	if err != nil {
		return errors.New("could not parse height")
	}

	if unit == Cm {
		if val < 150 || val > 193 {
			return errors.New("in, height too short or too long")
		}
	} else if unit == In {
		if val < 59 || val > 76 {
			return errors.New("cm: height too short or too long")
		}
	} else {
		// Invalid unit
		return errors.New("invalid height unit")
	}

	return nil
}

func validateHairColor(hairColor string) error {
	if hairColor[0] != '#' {
		return errors.New("hcl missing #")
	}

	if len(hairColor) != 7 {
		return errors.New("hcl too short")
	}

	for _, c := range hairColor[1:] {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return errors.New("hcl non-hex char")
		}
	}

	return nil
}

func parseInt(str string) (int, error) {
	v, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func mustParseInt(str string) int {
	v, err := parseInt(str)

	if err != nil {
		panic(err)
	}

	return v
}

func parseLines(lines []string) []passport {
	passports := make([]passport, 0)

	pp := passport{}

	for _, line := range lines {
		if line == "" {
			passports = append(passports, pp)
			pp = passport{}
			continue
		}

		statements := strings.Split(line, " ")

		for _, statement := range statements {
			pair := strings.Split(statement, ":")

			field := FieldType(pair[0])
			value := pair[1]

			switch field {
			case Byr:
				pp.byr = mustParseInt(value)
			case Iyr:
				pp.iyr = mustParseInt(value)
			case Eyr:
				pp.eyr = mustParseInt(value)
			case Hgt:
				pp.hgt = value
			case Hcl:
				pp.hcl = value
			case Ecl:
				pp.ecl = value
			case Pid:
				pp.pid = value
			case Cid:
				fallthrough
			default:
				// Ignore
				continue
			}
		}
	}

	return passports
}

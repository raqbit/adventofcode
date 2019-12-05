package intcode

import (
	"strconv"
	"strings"
)

// Parses the input lines as integers
func ParseIntcodes(input string) ([]int, error) {
	unparsedIntcodes := strings.Split(strings.TrimSpace(input), ",")
	intcodes := make([]int, len(unparsedIntcodes))
	for i, uic := range unparsedIntcodes {
		code, err := strconv.ParseInt(uic, 10, 32)
		if err != nil {
			return nil, err
		}
		intcodes[i] = int(code)
	}
	return intcodes, nil
}

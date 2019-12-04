package shared

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Taxicab distance
func (p Point) DistTo(other Point) int {
	dist := math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y))
	// Can safely cast since result of taxicab distance should never be a non-whole number
	return int(dist)
}

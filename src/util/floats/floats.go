package floats

import "math"

// AlmostEquals returns whether two floats are equal up to a small epsilon
func AlmostEquals(a, b float64) bool {
	return math.Abs(a-b) < math.SmallestNonzeroFloat64
}

// IsSmaller returns true iff a is smaller than b by (at least) epsilon
func IsSmaller(a, b float64) bool {
	return a+math.SmallestNonzeroFloat64 < b
}

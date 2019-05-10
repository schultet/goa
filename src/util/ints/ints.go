package ints

const (
	// MaxValue is the biggest value representable with type int
	MaxValue = int(^uint(0) >> 1)
	// MinValue is the smallest representable int value
	MinValue = -MaxValue
)

// Sum returns the sum of a list of ints
func Sum(vs ...int) int {
	sum := 0
	for _, val := range vs {
		sum += val
	}
	return sum
}

// Min returns the minimum value of a list of ints
func Min(vs ...int) int {
	min := MaxValue
	for _, val := range vs {
		if val < min {
			min = val
		}
	}
	return min
}

// Max returns the maximum value of a list of ints
func Max(vs ...int) int {
	max := MinValue
	for _, val := range vs {
		if val > max {
			max = val
		}
	}
	return max
}

// Mean returns the mean of a list of int values
func Mean(vs ...int) float64 {
	acc := 0.0
	for _, v := range vs {
		acc += float64(v)
	}
	return acc / float64(len(vs))
}

// Min2 returns the smaller of two ints
func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

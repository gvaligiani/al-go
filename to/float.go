package to

import (
	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// convert to float

func Float32[T constraints.Float](value T) float32 {
	return float32(value)
}

func Float64[T constraints.Float](value T) float64 {
	return float64(value)
}

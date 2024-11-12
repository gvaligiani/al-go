package to

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// convert to string

func Str[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

func IntStr[T constraints.Integer](value T) string {
	return fmt.Sprintf("%d", value)
}

func FloatStr[T constraints.Float](value T) string {
	return fmt.Sprintf("%f", value)
}

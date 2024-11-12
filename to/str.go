package to

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// convert to string

// Str convert simply the value into a string using fmt %v
func Str[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

// BiStr convert simply the pair into a string using fmt %v
func BiStr[K any, V any](key K, value V) string {
	return fmt.Sprintf("%v:%v", key, value)
}

func IntStr[T constraints.Integer](value T) string {
	return fmt.Sprintf("%d", value)
}

func FloatStr[T constraints.Float](value T) string {
	return fmt.Sprintf("%f", value)
}

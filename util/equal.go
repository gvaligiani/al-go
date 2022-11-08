package util

import "reflect"

func Equal[T comparable](left T, right T) bool {
	return left == right
}

func DeepEqual[T any](left T, right T) bool {
	return reflect.DeepEqual(left, right)
}

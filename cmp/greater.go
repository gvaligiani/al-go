package util

import "golang.org/x/exp/constraints"

func Greater[V constraints.Ordered](left V, right V) bool {
	return left > right
}

func GreaterOrEqual[V constraints.Ordered](left V, right V) bool {
	return left >= right
}

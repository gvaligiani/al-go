package util

import "golang.org/x/exp/constraints"

func Less[V constraints.Ordered](left V, right V) bool {
	return left < right
}

func LessOrEqual[V constraints.Ordered](left V, right V) bool {
	return left <= right
}

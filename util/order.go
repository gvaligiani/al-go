package util

import "golang.org/x/exp/constraints"

func Min[V constraints.Ordered](left V, right V) V {
	if left < right {
		return left
	}
	return right
}

func Max[V constraints.Ordered](left V, right V) V {
	if left < right {
		return right
	}
	return left
}

func Less[V constraints.Ordered](left V, right V) bool {
	return left < right
}

func LessOrEqual[V constraints.Ordered](left V, right V) bool {
	return left <= right
}

func Greater[V constraints.Ordered](left V, right V) bool {
	return left > right
}

func GreaterOrEqual[V constraints.Ordered](left V, right V) bool {
	return left >= right
}

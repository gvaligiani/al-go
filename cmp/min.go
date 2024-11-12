package util

import "golang.org/x/exp/constraints"

func Min[V constraints.Ordered](left V, right V) V {
	if left < right {
		return left
	}
	return right
}

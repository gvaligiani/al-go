package util

import "golang.org/x/exp/constraints"

func Max[V constraints.Ordered](left V, right V) V {
	if left < right {
		return right
	}
	return left
}

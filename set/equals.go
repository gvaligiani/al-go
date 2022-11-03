package set

func Equals[T comparable, S ~map[T]struct{}](left S, right S) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for leftItem := range left {
		// TODO not optimal >>> how to compare struct pointer
		if !Find(right, leftItem) {
			return false
		}
	}
	return true
}

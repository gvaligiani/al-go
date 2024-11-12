package set

// //////////////////////////////////////////////////
// diff

func Diff[T comparable](left, right map[T]struct{}) map[T]struct{} {
	diff := make(map[T]struct{}, len(left))
	for item := range left {
		if _, ok := right[item]; !ok {
			diff[item] = struct{}{}
		}
	}
	return diff
}

package set

// //////////////////////////////////////////////////
// intersect

func Intersect[T comparable](left, right map[T]struct{}) map[T]struct{} {
	intersect := make(map[T]struct{}, len(left))
	for item := range left {
		if _, ok := right[item]; ok {
			intersect[item] = struct{}{}
		}
	}
	return intersect
}

package set

// //////////////////////////////////////////////////
// union

func Union[T comparable](left, right map[T]struct{}) map[T]struct{} {
	union := make(map[T]struct{}, len(left)+len(right))
	for item := range left {
		union[item] = struct{}{}
	}
	for item := range right {
		union[item] = struct{}{}
	}
	return union
}

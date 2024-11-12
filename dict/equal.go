package dict

// //////////////////////////////////////////////////
// equal

func Equal[K comparable, V comparable](left, right map[K]V) bool {
	if len(left) != len(right) {
		return false
	}
	for key, value := range left {
		if value != right[key] {
			return false
		}
	}
	return true
}

package set

// //////////////////////////////////////////////////
// to list

// ToList converts the set into a list
func ToList[T comparable](items map[T]struct{}) []T {
	converted := make([]T, 0, len(items))
	for item := range items {
		converted = append(converted, item)
	}
	return converted
}

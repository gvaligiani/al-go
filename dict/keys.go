package dict

// //////////////////////////////////////////////////
// keys

// Keys extracts map keys as a slice
func Keys[K comparable, V any](items map[K]V) []K {
	converted := make([]K, 0, len(items))
	for key := range items {
		converted = append(converted, key)
	}
	return converted
}

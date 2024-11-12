package dict

// //////////////////////////////////////////////////
// to list

func ToList[K comparable, V any](items map[K]V) []V {
	converted := make([]V, 0, len(items))
	for _, item := range items {
		converted = append(converted, item)
	}
	return converted
}

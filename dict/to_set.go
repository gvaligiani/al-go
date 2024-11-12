package dict

// //////////////////////////////////////////////////
// to set

func ToSet[K comparable, V comparable](items map[K]V) map[V]struct{} {
	converted := make(map[V]struct{}, len(items))
	for _, item := range items {
		converted[item] = struct{}{}
	}
	return converted
}

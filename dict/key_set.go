package dict

// //////////////////////////////////////////////////
// key set

// KeySet extracts map keys as a set
func KeySet[K comparable, V any](items map[K]V) map[K]struct{} {
	keys := make(map[K]struct{}, len(items))
	for key := range items {
		keys[key] = struct{}{}
	}
	return keys
}

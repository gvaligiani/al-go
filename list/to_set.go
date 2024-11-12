package list

// //////////////////////////////////////////////////
// to set

func ToSet[T comparable](items []T) map[T]struct{} {
	converted := make(map[T]struct{}, len(items))
	for _, item := range items {
		converted[item] = struct{}{}
	}
	return converted
}

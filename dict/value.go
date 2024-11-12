package dict

// /////////////////////////////////////////////////
// value helper

func Value[K comparable, V any](k K, v V) V {
	return v
}

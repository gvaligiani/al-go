package dict

// /////////////////////////////////////////////////
// key helper

func Key[K comparable, V any](k K, v V) K {
	return k
}

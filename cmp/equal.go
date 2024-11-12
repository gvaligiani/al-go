package util

func Equal[V comparable](left V, right V) bool {
	return left == right
}

func NotEqual[V comparable](left V, right V) bool {
	return left != right
}

package util

import "reflect"

func Equal[V comparable](left V, right V) bool {
	return left == right
}

func NotEqual[V comparable](left V, right V) bool {
	return left != right
}

func DeepEqual[V any](left V, right V) bool {
	return reflect.DeepEqual(left, right)
}

func DeepNotEqual[V any](left V, right V) bool {
	return !reflect.DeepEqual(left, right)
}

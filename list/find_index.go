package list

import (
	"github.com/gvaligiani/al.go/util"
)

func FindIndex[T comparable, L ~[]T](items L, index int) bool {
	if 0 <= index && index < len(items) {
		return true
	}
	return false
}

func DeepFindIndex[T any, L ~[]T](items L, index int) bool {
	return FindIndexFn(items, index, util.DeepEqual[int])
}

func FindIndexFn[T any, L ~[]T](items L, index int, equal util.BiPredicate[int, int]) bool {
	_, _, found := FindIndexIf(items, func(i int, _ T) bool { return equal(i, index) })
	return found
}

func FindValueAt[T comparable, L ~[]T](items L, index int) (T, bool) {
	if 0 <= index && index < len(items) {
		return items[index], true
	}
	var noValue T
	return noValue, false
}

func DeepFindValueAt[T any, L ~[]T](items L, index int) (T, bool) {
	return FindValueAtFn(items, index, util.DeepEqual[int])
}

func FindValueAtFn[T any, L ~[]T](items L, index int, equal util.BiPredicate[int, int]) (T, bool) {
	_, value, found := FindIndexIf(items, func(i int, _ T) bool { return equal(i, index) })
	return value, found
}

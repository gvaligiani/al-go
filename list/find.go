package list

import "reflect"

func Find[T any, L ~[]T](items L, value T) (int, bool) {
	index, _, found := FindIf(items, func(_ int, item T) bool { return reflect.DeepEqual(item, value) })
	return index, found
}

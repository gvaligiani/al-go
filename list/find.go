package list

import "reflect"

func Find[T any, L ~[]T](items L, value T) bool {
	_, found := FindIf(items, func(item T) bool { return reflect.DeepEqual(item, value) })
	return found
}

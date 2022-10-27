package list

import "reflect"

func Find[T comparable](items []T, value T) bool {
	_, found := FindIf(items, func(item T) bool { return reflect.DeepEqual(item,value) })
	return found
}
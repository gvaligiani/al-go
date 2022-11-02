package list

import "reflect"

func Find[T any](items []T, value T) bool {
	_, found := FindIf(items, func(item T) bool { return reflect.DeepEqual(item,value) })
	return found
}
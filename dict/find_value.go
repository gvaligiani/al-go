package dict

import "reflect"

func FindValue[K comparable, T any, M ~map[K]T](items M, value T) bool {
	_, _, found := FindIf(items, func(_ K, item T) bool { return reflect.DeepEqual(item, value) })
	return found
}

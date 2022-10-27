package dict

import "reflect"

func Find[K comparable,T comparable](items map[K]T, value T) bool {
	_, found := FindIf(items, func(_ K, item T) bool { return reflect.DeepEqual(item,value) })
	return found
}
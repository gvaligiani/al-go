package dict

type Dict[K comparable,T any] map[K]T

// builder

func New[K comparable,T any]() Dict[K,T] {
	return Dict[K,T]{}
}

// helper

func (d Dict[K,T]) Add(key K, item T) bool {
	_, overriden := d[key]
	d[key] = item
	return overriden
}

func (d Dict[K,T]) Remove(key K) bool {
	if _, ok := d[key]; ok {
		delete(d, key)
		return true
	}
	return false
}

func (d *Dict[K,T]) Clear() {
	*d = map[K]T{}
}

func (d Dict[K,T]) Keys() []K {
	l := make([]K,0,len(d))
	for key := range d {
		l = append(l,key)
	}
	return l
}

func (d Dict[K,T]) Values() []T {
	l := make([]T,0,len(d))
	for _, item := range d {
		l = append(l,item)
	}
	return l
}

// algo

func (d Dict[K,T]) AllOf(predicate func(K,T) bool) bool {
	return AllOf[K,T](d,predicate)
}

func (d Dict[K,T]) AnyOf(predicate func(K,T) bool) bool {
	return AnyOf[K,T](d,predicate)
}

func (d Dict[K,T]) NoneOf(predicate func(K,T) bool) bool {
	return NoneOf[K,T](d,predicate)
}

func (d Dict[K,T]) Each(consumer func(K,T)) {
	Each[K,T](d,consumer)
}

func (d Dict[K,T]) FindKey(key K) bool {
	return FindKey[K,T](d,key)
}

func (d Dict[K,T]) FindValue(value T) bool {
	return FindValue[K,T](d,value)
}

func (d Dict[K,T]) FindIf(predicate func(K,T) bool) (K,T,bool) {
	return FindIf[K,T](d,predicate)
}

func (d Dict[K,T]) FindIfNot(predicate func(K,T) bool) (K,T,bool) {
	return FindIfNot[K,T](d,predicate)
}
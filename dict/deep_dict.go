package dict

// alias

type DeepDict[K comparable, T any] map[K]T

// builder

func NewDeep[K comparable, T any]() DeepDict[K, T] {
	return DeepDict[K, T]{}
}

// getter

func (d DeepDict[K, T]) Keys() []K {
	l := make([]K, 0, len(d))
	for key := range d {
		l = append(l, key)
	}
	return l
}

func (d DeepDict[K, T]) Values() []T {
	l := make([]T, 0, len(d))
	for _, item := range d {
		l = append(l, item)
	}
	return l
}

// state

func (d DeepDict[K, T]) IsEmpty() bool {
	return len(d) == 0
}

func (d DeepDict[K, T]) AllOf(predicate Predicate[K, T]) bool {
	return AllOf(d, predicate)
}

func (d DeepDict[K, T]) AnyOf(predicate Predicate[K, T]) bool {
	return AnyOf(d, predicate)
}

func (d DeepDict[K, T]) NoneOf(predicate Predicate[K, T]) bool {
	return NoneOf(d, predicate)
}

// each

func (d DeepDict[K, T]) Each(consumer Consumer[K, T]) {
	Each(d, consumer)
}

// find

func (d DeepDict[K, T]) FindKey(key K) bool {
	return FindKey(d, key)
}

func (d DeepDict[K, T]) FindValue(value T) bool {
	return DeepFindValue(d, value)
}

func (d DeepDict[K, T]) FindIf(predicate Predicate[K, T]) (K, T, bool) {
	return FindIf(d, predicate)
}

func (d DeepDict[K, T]) FindIfNot(predicate Predicate[K, T]) (K, T, bool) {
	return FindIfNot(d, predicate)
}

// copy

func (d DeepDict[K, T]) Copy() DeepDict[K, T] {
	return DeepDict[K, T](Copy(d))
}

func (d DeepDict[K, T]) CopyIf(predicate Predicate[K, T]) DeepDict[K, T] {
	return DeepDict[K, T](CopyIf(d, predicate))
}

func (d DeepDict[K, T]) CopyIfNot(predicate Predicate[K, T]) DeepDict[K, T] {
	return DeepDict[K, T](CopyIfNot(d, predicate))
}

// modifier

func (d *DeepDict[K, T]) Add(key K, item T) bool {
	_, overriden := (*d)[key]
	(*d)[key] = item
	return overriden
}

func (d *DeepDict[K, T]) Remove(key K) bool {
	if _, ok := (*d)[key]; ok {
		delete(*d, key)
		return true
	}
	return false
}

func (d *DeepDict[K, T]) Clear() bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	*d = DeepDict[K, T]{}
	return true
}

func (d *DeepDict[K, T]) RemoveIf(predicate Predicate[K, T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return RemoveIf(d, predicate)
}

func (d *DeepDict[K, T]) KeepIf(predicate Predicate[K, T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return KeepIf(d, predicate)
}

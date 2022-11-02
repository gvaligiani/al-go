package list

import "reflect"

type List[T any] []T

// builder

func New[T comparable](values ...T) List[T] {
	l := make(List[T],0,len(values))
	l = append(l, values...)
	return l
}

// helper

// TODO why *
func (l *List[T]) Add(values ...T) bool {
	*l = append(*l, values...)
	return true
}

func (l *List[T]) Remove(value T) bool {
	newL := make([]T,0,len(*l))
	atLeastOneRemoved := false
	for _, item := range *l {
		if reflect.DeepEqual(item,value) {
			atLeastOneRemoved = true
		} else {
			newL = append( newL, item )
		}
	}
	*l = newL
	return atLeastOneRemoved
}

func (l *List[T]) Clear() {
	*l = []T{}
}

// algo

func (l List[T]) AllOf(predicate func(T) bool) bool {
	return AllOf[T](l,predicate)
}

func (l List[T]) AnyOf(predicate func(T) bool) bool {
	return AnyOf[T](l,predicate)
}

func (l List[T]) NoneOf(predicate func(T) bool) bool {
	return NoneOf[T](l,predicate)
}

func (l List[T]) Each(consumer func(T)) {
	Each[T](l,consumer)
}

func (l List[T]) EachIndex(consumer func(int,T)) {
	EachIndex[T](l,consumer)
}

func (l List[T]) Find(value T) bool {
	return Find[T](l,value)
}

func (l List[T]) FindIf(predicate func(T) bool) (T,bool) {
	return FindIf[T](l,predicate)
}

func (l List[T]) FindIfNot(predicate func(T) bool) (T,bool) {
	return FindIfNot[T](l,predicate)
}
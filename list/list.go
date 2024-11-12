package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// list

type List[T any] []T

// //////////////////////////////////////////////////
// add

// Add allows to add multiple values in the list
func (l *List[T]) Add(values ...T) {
	(*l) = append((*l), values...)
}

// //////////////////////////////////////////////////
// state

func (l List[T]) AllOf(predicate fn.Predicate[T]) bool {
	return AllOf(l, predicate)
}

func (l List[T]) AnyOf(predicate fn.Predicate[T]) bool {
	return AnyOf(l, predicate)
}

func (l List[T]) NoneOf(predicate fn.Predicate[T]) bool {
	return NoneOf(l, predicate)
}

// //////////////////////////////////////////////////
// remove if

func (l *List[T]) RemoveIf(predicate fn.Predicate[T]) {
	RemoveIf(l, predicate)
}

// //////////////////////////////////////////////////
// keep if

func (l *List[T]) KeepIf(predicate fn.Predicate[T]) {
	KeepIf(l, predicate)
}

// //////////////////////////////////////////////////
// count if

func (l List[T]) CountIf(predicate fn.Predicate[T]) int {
	return CountIf(l, predicate)
}

// //////////////////////////////////////////////////
// each

func (l List[T]) Each(consume fn.Consumer[T]) {
	Each(l, consume)
}

// //////////////////////////////////////////////////
// copy if

func (l List[T]) CopyIf(predicate fn.Predicate[T]) List[T] {
	return List[T](CopyIf(l, predicate))
}

// //////////////////////////////////////////////////
// find

func (l List[T]) FindIf(predicate fn.Predicate[T]) (T, bool) {
	return FindIf(l, predicate)
}

// //////////////////////////////////////////////////
// join

func (l List[T]) Join(separator string) string {
	return Join(l, separator)
}

func (l List[T]) ConvertAndJoin(convert fn.Converter[T, string], separator string) string {
	return ConvertAndJoin(l, convert, separator)
}

// //////////////////////////////////////////////////
// validate

func (l List[T]) Validate(validate fn.Validator[T]) error {
	return Validate(l, validate)
}

package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// set

type Set[T comparable] map[T]struct{}

// //////////////////////////////////////////////////
// has

// Has is an helper to check if a vaue is present in the set
func (s Set[T]) Has(value T) bool {
	_, found := s[value]
	return found
}

// //////////////////////////////////////////////////
// add

// Add is an helper to add multiple values in the set
func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		(*s)[value] = struct{}{}
	}
}

// //////////////////////////////////////////////////
// del

// Del is an helper to remove multiple values from the set
func (s *Set[T]) Del(values ...T) {
	for _, value := range values {
		delete((*s), value)
	}
}

// //////////////////////////////////////////////////
// to list

func (s Set[T]) ToList(values ...T) []T {
	return ToList(s)
}

// //////////////////////////////////////////////////
// state

func (s Set[T]) AllOf(predicate fn.Predicate[T]) bool {
	return AllOf(s, predicate)
}

func (s Set[T]) AnyOf(predicate fn.Predicate[T]) bool {
	return AnyOf(s, predicate)
}

func (s Set[T]) NoneOf(predicate fn.Predicate[T]) bool {
	return NoneOf(s, predicate)
}

// //////////////////////////////////////////////////
// remove if

func (s *Set[T]) RemoveIf(predicate fn.Predicate[T]) {
	RemoveIf(s, predicate)
}

// //////////////////////////////////////////////////
// keep if

func (s *Set[T]) KeepIf(predicate fn.Predicate[T]) {
	KeepIf(s, predicate)
}

// //////////////////////////////////////////////////
// count if

func (s Set[T]) CountIf(predicate fn.Predicate[T]) int {
	return CountIf(s, predicate)
}

// //////////////////////////////////////////////////
// each

func (s Set[T]) Each(consume fn.Consumer[T]) {
	Each(s, consume)
}

// //////////////////////////////////////////////////
// copy if

func (s Set[T]) CopyIf(predicate fn.Predicate[T]) Set[T] {
	return Set[T](CopyIf(s, predicate))
}

// //////////////////////////////////////////////////
// find

func (s Set[T]) FindIf(predicate fn.Predicate[T]) (T, bool) {
	return FindIf(s, predicate)
}

// //////////////////////////////////////////////////
// join

func (s Set[T]) Join(separator string) string {
	return Join(s, separator)
}

func (s Set[T]) ConvertAndJoin(convert fn.Converter[T, string], separator string) string {
	return ConvertAndJoin(s, convert, separator)
}

// //////////////////////////////////////////////////
// validate

func (s Set[T]) Validate(validate fn.Validator[T]) error {
	return Validate(s, validate)
}

// //////////////////////////////////////////////////
// diff

func (s Set[T]) Diff(o Set[T]) Set[T] {
	return Set[T](Diff(s, o))
}

// //////////////////////////////////////////////////
// intersect

func (s Set[T]) Intersect(o Set[T]) Set[T] {
	return Set[T](Intersect(s, o))
}

// //////////////////////////////////////////////////
// union

func (s Set[T]) Union(o Set[T]) Set[T] {
	return Set[T](Union(s, o))
}

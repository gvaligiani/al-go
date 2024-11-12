package set

import (
	"strings"

	"github.com/gvaligiani/al-go/fn"
	"github.com/gvaligiani/al-go/to"
)

// //////////////////////////////////////////////////
// join

// Join converts items into string using fmt %v and joins them with the separator
func Join[T comparable](items map[T]struct{}, separator string) string {
	return ConvertAndJoin(items, to.Str, separator)
}

// ConvertAndJoin converts items into string using the provided converter and joins them with the separator
func ConvertAndJoin[T comparable](items map[T]struct{}, convert fn.Converter[T, string], separator string) string {
	return strings.Join(ConvertToList(items, convert), separator)
}

package list

import (
	"strings"

	"github.com/gvaligiani/al-go/fn"
	"github.com/gvaligiani/al-go/to"
)

// //////////////////////////////////////////////////
// join

// Join converts items into string using fmt %v and joins them with the separator
func Join[T any](items []T, separator string) string {
	return ConvertAndJoin(items, to.Str, separator)
}

// ConvertAndJoin converts items into string using the provided converter and joins them with the separator
func ConvertAndJoin[T any](items []T, convert fn.Converter[T, string], separator string) string {
	return strings.Join(Convert(items, convert), separator)
}

package dict

import (
	"strings"

	"github.com/gvaligiani/al-go/fn"
	"github.com/gvaligiani/al-go/to"
)

// //////////////////////////////////////////////////
// join

// Join converts key-value pairs into string using fmt %v and joins them with the separator
func Join[K comparable, V any](items map[K]V, separator string) string {
	return ConvertAndJoin(items, to.BiStr, separator)
}

// ConvertAndJoin converts items into string using the provided converter and joins them with the separator
func ConvertAndJoin[K comparable, V any](items map[K]V, convert fn.BiConverter[K, V, string], separator string) string {
	return strings.Join(ConvertToList(items, convert), separator)
}

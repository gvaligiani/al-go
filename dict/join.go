package dict

import (
	"fmt"
	"strings"

	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// join

// Join converts key-value pairs into string using fmt %v and joins them with the separator
func Join[K comparable, V any](items map[K]V, separator string) string {
	return ConvertAndJoin(items, ToStr, separator)
}

// ConvertAndJoin converts items into string using the provided converter and joins them with the separator
func ConvertAndJoin[K comparable, V any](items map[K]V, convert fn.BiConverter[K, V, string], separator string) string {
	return strings.Join(ConvertToList(items, convert), separator)
}

// ToStr convert simply the pair into a string using fmt %v
func ToStr[K any, V any](key K, value V) string {
	return fmt.Sprintf("%v:%v", key, value)
}

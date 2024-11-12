package set

import (
	"fmt"
	"strings"

	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// join

// Join converts items into string using fmt %v and joins them with the separator
func Join[T comparable](items map[T]struct{}, separator string) string {
	return ConvertAndJoin(items, ToStr, separator)
}

// ConvertAndJoin converts items into string using the provided converter and joins them with the separator
func ConvertAndJoin[T comparable](items map[T]struct{}, convert fn.Converter[T, string], separator string) string {
	return strings.Join(ConvertToList(items, convert), separator)
}

// ToStr convert simply the item into a string using fmt %v
func ToStr[T any](item T) string {
	return fmt.Sprintf("%v", item)
}

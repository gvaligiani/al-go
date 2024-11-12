package list

import (
	"fmt"
	"strings"

	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// join

// Join converts items into string using fmt %v and joins them with the separator
func Join[T any](items []T, separator string) string {
	return ConvertAndJoin(items, ToStr, separator)
}

// ConvertAndJoin converts items into string using the provided converter and joins them with the separator
func ConvertAndJoin[T any](items []T, convert fn.Converter[T, string], separator string) string {
	return strings.Join(Convert(items, convert), separator)
}

// ToStr convert simply the item into a string using fmt %v
func ToStr[T any](item T) string {
	return fmt.Sprintf("%v", item)
}

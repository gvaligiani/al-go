package set

import "github.com/gvaligiani/algo/dict"

func Find[T comparable](items map[T]struct{}, value T) bool {
	return dict.FindKey[T,struct{}](items,value)
}
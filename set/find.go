package set

import "github.com/gvaligiani/al.go/dict"

func Find[T comparable, S ~map[T]struct{}](items S, value T) bool {
	return dict.FindKey(items, value)
}

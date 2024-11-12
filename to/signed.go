package to

import (
	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// convert to signed

func Int[T constraints.Signed](value T) int {
	return int(value)
}

func Int8[T constraints.Signed](value T) int8 {
	return int8(value)
}

func Int16[T constraints.Signed](value T) int16 {
	return int16(value)
}

func Int32[T constraints.Signed](value T) int32 {
	return int32(value)
}

func Int64[T constraints.Signed](value T) int64 {
	return int64(value)
}

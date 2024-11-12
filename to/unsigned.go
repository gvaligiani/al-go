package to

import (
	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// convert to unsigned

func UInt[T constraints.Unsigned](value T) uint {
	return uint(value)
}

func UInt8[T constraints.Unsigned](value T) uint8 {
	return uint8(value)
}

func UInt16[T constraints.Unsigned](value T) uint16 {
	return uint16(value)
}

func UInt32[T constraints.Unsigned](value T) uint32 {
	return uint32(value)
}

func UInt64[T constraints.Unsigned](value T) uint64 {
	return uint64(value)
}

func UIntPtr[T constraints.Unsigned](value T) uintptr {
	return uintptr(value)
}

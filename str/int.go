package str

import (
	"strconv"
)

// //////////////////////////////////////////////////
// convert string to

func ToInt(value string) int {
	if value, err := strconv.Atoi(value); err == nil {
		return value
	}
	return 0
}

func ToInt64(value string) int64 {
	if value, err := strconv.ParseInt(value, 10, 64); err == nil {
		return value
	}
	return 0
}

package list_test

import (
	"fmt"
	"strconv"

	"github.com/gvaligiani/al-go/list"
)

func IsEven(v int) bool {
	return v%2 == 0
}

type A struct {
	value int
}

func (a *A) IsEven() bool {
	return IsEven(a.value)
}

func (a *A) Validate() error {
	if a.value < 0 {
		return fmt.Errorf("negative")
	}
	if a.value > 15 {
		return fmt.Errorf("too big")
	}
	return nil
}

func Example() {
	ints := []int{2, 6, 10, 17, 16}

	// convert
	converted := list.Convert(ints, strconv.Itoa)
	fmt.Printf("%v\n", converted)

	// decode
	decoded, err := list.StrictDecode(ints, func(i int) (*A, error) {
		if i < 0 {
			return nil, fmt.Errorf("negative")
		}
		return &A{value: i}, nil
	})
	fmt.Printf("%v\n", err)

	// all of
	notAboveTwenty := list.AllOf(ints, func(i int) bool { return i < 20 })
	fmt.Printf("%t\n", notAboveTwenty)

	// any of
	atLeastOneEven := list.AnyOf(decoded, func(a *A) bool { return a.value%2 == 0 })
	fmt.Printf("%t\n", atLeastOneEven)

	// validate
	err = list.Validate(decoded, (*A).Validate)
	fmt.Printf("%v\n", err)

	// Output:
	// [2 6 10 17 16]
	// <nil>
	// true
	// true
	// too big
}

func ExampleAllOf() {
	values := []int{2, 6, 10, 17, 16}
	fmt.Printf("%t\n", list.AllOf(values, func(v int) bool { return v < 20 }))
	fmt.Printf("%t\n", list.AllOf(values, IsEven))
	// Output:
	// true
	// false
}

func ExampleAnyOf() {
	values := []int{2, 6, 10, 17, 16}
	fmt.Printf("%t\n", list.AnyOf(values, func(v int) bool { return v > 20 }))
	fmt.Printf("%t\n", list.AnyOf(values, IsEven))
	// Output:
	// false
	// true
}

func ExampleNoneOf() {
	values := []int{2, 6, 10, 17, 16}
	fmt.Printf("%t\n", list.NoneOf(values, func(v int) bool { return v > 20 }))
	fmt.Printf("%t\n", list.NoneOf(values, IsEven))
	// Output:
	// true
	// false
}

func ExampleConvert() {
	values := []int{2, 6, 10, 17, 16}
	fmt.Printf("%t\n", list.Convert(values, func(v int) bool { return v%2 == 0 }))
	// Output:
	// [true true true false true]
}

func ExampleStrictDecode_success() {
	values := []int{2, 6, 10, 16}
	decoded, err := list.StrictDecode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, fmt.Errorf("%d is not even", v)
		}
		return v / 2, nil
	})
	fmt.Printf("%v\n", decoded)
	fmt.Printf("%v\n", err)
	// Output:
	// [1 3 5 8]
	// <nil>
}

func ExampleStrictDecode_error() {
	values := []int{2, 6, 10, 17, 16}
	decoded, err := list.StrictDecode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, fmt.Errorf("%d is not even", v)
		}
		return v / 2, nil
	})
	fmt.Printf("%v\n", decoded)
	fmt.Printf("%v\n", err)
	// Output:
	// []
	// 17 is not even
}

func ExampleDecode_success() {
	values := []int{2, 6, 10, 16}
	decoded := list.Decode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, fmt.Errorf("not even")
		}
		return v / 2, nil
	}, func(v int, err error) {
		fmt.Printf("%v\n", err)
	})
	fmt.Printf("%v\n", decoded)
	// Output:
	// [1 3 5 8]
}

func ExampleDecode_error() {
	values := []int{2, 6, 10, 17, 16}
	decoded := list.Decode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, fmt.Errorf("%d is not even", v)
		}
		return v / 2, nil
	}, func(v int, err error) {
		fmt.Printf("%v\n", err)
	})
	fmt.Printf("%v\n", decoded)
	// Output:
	// 17 is not even
	// [1 3 5 8]
}

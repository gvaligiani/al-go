package set_test

import (
	"fmt"
	"strconv"

	"github.com/gvaligiani/al-go/set"
	"github.com/travelaudience/cxt-go/testx"
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
	ints := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}

	// convert
	converted := set.Convert(ints, strconv.Itoa)
	fmt.Printf("%v\n", converted)

	// decode
	decoded, err := set.StrictDecode(ints, func(i int) (*A, error) {
		if i < 0 {
			return nil, fmt.Errorf("negative")
		}
		return &A{value: i}, nil
	})
	fmt.Printf("%v\n", err)

	// all of
	notAboveTwenty := set.AllOf(ints, func(i int) bool { return i < 20 })
	fmt.Printf("%t\n", notAboveTwenty)

	// any of
	atLeastOneEven := set.AnyOf(decoded, func(a *A) bool { return a.value%2 == 0 })
	fmt.Printf("%t\n", atLeastOneEven)

	// validate
	err = set.Validate(decoded, (*A).Validate)
	fmt.Printf("%v\n", err)

	// Output:
	// map[10:{} 16:{} 17:{} 2:{} 6:{}]
	// <nil>
	// true
	// true
	// too big
}

func ExampleAllOf() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}
	fmt.Printf("%t\n", set.AllOf(values, func(v int) bool { return v < 20 }))
	fmt.Printf("%t\n", set.AllOf(values, IsEven))
	// Output:
	// true
	// false
}

func ExampleAnyOf() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}
	fmt.Printf("%t\n", set.AnyOf(values, func(v int) bool { return v > 20 }))
	fmt.Printf("%t\n", set.AnyOf(values, IsEven))
	// Output:
	// false
	// true
}

func ExampleNoneOf() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}
	fmt.Printf("%t\n", set.NoneOf(values, func(v int) bool { return v > 20 }))
	fmt.Printf("%t\n", set.NoneOf(values, IsEven))
	// Output:
	// true
	// false
}

func ExampleConvert() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}
	fmt.Printf("%t\n", set.Convert(values, func(v int) bool { return v%2 == 0 }))
	// Output:
	// map[false:{} true:{}]
}

func ExampleStrictDecode_success() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 16: {}}
	decoded, err := set.StrictDecode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, testx.Error
		}
		return v / 2, nil
	})
	fmt.Printf("%v\n", decoded)
	fmt.Printf("%v\n", err)
	// Output:
	// map[1:{} 3:{} 5:{} 8:{}]
	// <nil>
}

func ExampleStrictDecode_error() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}
	decoded, err := set.StrictDecode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, fmt.Errorf("%d is not even", v)
		}
		return v / 2, nil
	})
	fmt.Printf("%v\n", decoded)
	fmt.Printf("%v\n", err)
	// Output:
	// map[]
	// fmt.Errorf("%d is not even", v)
}

func ExampleDecode_success() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 16: {}}
	decoded := set.Decode(values, func(v int) (int, error) {
		if v%2 != 0 {
			return 0, fmt.Errorf("%d is not even", v)
		}
		return v / 2, nil
	}, func(v int, err error) {
		fmt.Printf("%v\n", err)
	})
	fmt.Printf("%v\n", decoded)
	// Output:
	// map[1:{} 3:{} 5:{} 8:{}]
}

func ExampleDecode_error() {
	values := map[int]struct{}{2: {}, 6: {}, 10: {}, 17: {}, 16: {}}
	decoded := set.Decode(values, func(v int) (int, error) {
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
	// map[1:{} 3:{} 5:{} 8:{}]
}

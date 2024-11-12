package set

import "github.com/gvaligiani/al-go/fn"

func Validate[T comparable](items map[T]struct{}, validate fn.Validator[T]) error {
	for item := range items {
		if err := validate(item); err != nil {
			return err
		}
	}
	return nil
}

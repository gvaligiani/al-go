package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// validate

func Validate[T any](items []T, validate fn.Validator[T]) error {
	for _, item := range items {
		if err := validate(item); err != nil {
			return err
		}
	}
	return nil
}

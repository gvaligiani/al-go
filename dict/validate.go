package dict

import "github.com/gvaligiani/al-go/fn"

func Validate[K comparable, V any](items map[K]V, validate fn.BiValidator[K, V]) error {
	for key, value := range items {
		if err := validate(key, value); err != nil {
			return err
		}
	}
	return nil
}

package dict

func Ptr[K comparable, V any, D ~map[K]V](d D) *D {
	if d == nil {
		return nil
	}
	return &d
}

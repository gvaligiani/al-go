package set

func Ptr[V comparable, S ~map[V]struct{}](s S) *S {
	if s == nil {
		return nil
	}
	return &s
}

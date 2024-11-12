package list

func Ptr[V any, L ~[]V](l L) *L {
	if l == nil {
		return nil
	}
	return &l
}

package list

// //////////////////////////////////////////////////
// new

func New[T any](items ...T) List[T] {
	list := make(List[T], 0, len(items))
	list = append(list, items...)
	return list
}

package set

type Set[T comparable] map[T]struct{}

// builder

func New[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	for _, item := range items {
		s.Add(item)
	}	
	return s
}

// helper

func (s Set[T]) Add(item T) bool {
	if _, ok := s[item]; !ok {
		s[item] = struct{}{}
		return true
	}
	return false
}

func (s Set[T]) Remove(item T) bool {
	if _, ok := s[item]; ok {
		delete(s, item)
		return true
	}
	return false
}

func (s *Set[T]) Clear() {
	*s = map[T]struct{}{}
}

func (s Set[T]) Values() []T {
	l := make([]T,0,len(s))
	for item := range s {
		l = append(l,item)
	}
	return l
}

// algo

func (s Set[T]) AllOf(predicate func(T) bool) bool {
	return AllOf[T](s,predicate)
}

func (s Set[T]) AnyOf(predicate func(T) bool) bool {
	return AnyOf[T](s,predicate)
}

func (s Set[T]) NoneOf(predicate func(T) bool) bool {
	return NoneOf[T](s,predicate)
}

func (s Set[T]) Each(consumer func(T)) {
	Each[T](s,consumer)
}

func (s Set[T]) Find(value T) bool {
	return Find[T](s,value)
}

func (s Set[T]) FindIf(predicate func(T) bool) (T,bool) {
	return FindIf[T](s,predicate)
}

func (s Set[T]) FindIfNot(predicate func(T) bool) (T,bool) {
	return FindIfNot[T](s,predicate)
}
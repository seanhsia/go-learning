package common

type Set[T comparable] struct {
	container map[T]struct{}
}

func ToSet[T comparable](slice []T) *Set[T] {
	s := &Set[T]{}
	s.New()
	for _, ele := range slice {
		s.Add(ele)
	}
	return s
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.container))
	for ele := range s.container {
		slice = append(slice, ele)
	}
	return slice
}

func (s *Set[T]) New() {
	s.container = make(map[T]struct{})
}

func (s *Set[T]) Add(ele T) {
	if _, exist := s.container[ele]; !exist {
		s.container[ele] = struct{}{}
	}
}

func (s *Set[T]) Remove(ele T) {
	delete(s.container, ele)
}

func (s *Set[T]) Exists(ele T) bool {
	_, exist := s.container[ele]
	return exist
}

func (s1 *Set[T]) Intersect(s2 *Set[T]) {
	for ele := range s1.container {
		if !s2.Exists(ele) {
			s1.Remove(ele)
		}
	}
}

func (s1 *Set[T]) Union(s2 *Set[T]) {
	for ele, v := range s2.container {
		s1.container[ele] = v
	}
}

func (s1 *Set[T]) Diff(s2 *Set[T]) {
	for ele := range s1.container {
		if s2.Exists(ele) {
			s1.Remove(ele)
		}
	}
}

func (s1 *Set[T]) Equal(s2 *Set[T]) bool {
	for ele := range s1.container {
		if !s2.Exists(ele) {
			return false
		}
	}

	for ele := range s2.container {
		if !s1.Exists(ele) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Len() int {
	return len(s.container)
}

package common

func Filter[T any](s []T, f func(T) bool) []T {
	s2 := []T{}

	for _, ele := range s {
		if f(ele) {
			s2 = append(s2, ele)
		}
	}
	return s2
}

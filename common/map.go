package common

func Map[T1 any, T2 any](s []T1, f func(T1) T2) []T2 {
	s2 := make([]T2, 0, len(s))

	for _, ele := range s {
		s2 = append(s2, f(ele))
	}
	return s2
}

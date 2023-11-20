package common

func Reduce[T1, T2 any](s []T1, accumulator T2, f func(T2, T1) T2) T2 {
	a := accumulator
	for _, ele := range s {
		a = f(a, ele)
	}
	return a
}

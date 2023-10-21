package sliceutil

func Map[T, R any](s []T, mapper func(T) R) []R {
	r := make([]R, 0)

	for _, v := range s {
		r = append(r, mapper(v))
	}

	return r
}

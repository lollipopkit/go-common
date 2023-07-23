package util

func Contains[T comparable](s []T, e T) bool {
	for idx := range s {
		if s[idx] == e {
			return true
		}
	}
	return false
}

func Clear[T any](s *[]T) {
	*s = (*s)[:0]
}


package util

func Contains[T string | int | int64 | float64](s []T, e T) bool {
	for idx := range s {
		if s[idx] == e {
			return true
		}
	}
	return false
}
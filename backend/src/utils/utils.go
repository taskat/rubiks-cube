package utils

func Contains[T comparable](arr []T, elem T) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}

package utils

func Contains[T comparable](arr []T, elem T) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}

func Filter[T any](arr []T, predicate func(T) bool) []T {
	filtered := make([]T, 0)
	for _, elem := range arr {
		if predicate(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

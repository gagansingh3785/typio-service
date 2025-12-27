package utils

func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}

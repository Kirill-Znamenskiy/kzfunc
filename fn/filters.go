package fn

// IsZero returns true when the provided value is the zero value for its type.
func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}

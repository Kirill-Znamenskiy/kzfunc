package fn

func ConvertValToPtr[T any](val T) *T {
	return &val
}

func ConvertPtrToVal[T any](ptr *T) T {
	if ptr == nil {
		var z T
		return z
	}
	return *ptr
}

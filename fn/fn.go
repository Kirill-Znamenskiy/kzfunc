package fn

func Map[A, B any](items []A, f func(A) B) (ret []B) {
	ret = make([]B, 0, len(items))
	for _, item := range items {
		tmp := f(item)
		ret = append(ret, tmp)
	}
	return ret
}

func Filter[A comparable](items []A, f func(A) bool) (ret []A) {
	ret = make([]A, 0, len(items))
	for _, item := range items {
		if f(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func MapFilter[A, B any](items []A, f func(A) *B) (ret []B) {
	ret = make([]B, 0, len(items))
	for _, item := range items {
		tmp := f(item)
		if tmp != nil {
			ret = append(ret, *tmp)
		}
	}
	return ret
}

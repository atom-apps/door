package common

func PtrToValue[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func OneOf[T any](ptr ...*T) *T {
	for _, p := range ptr {
		if p != nil {
			return p
		}
	}
	return nil
}

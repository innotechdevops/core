package pointer

func New[T any](value T) *T {
	return &value
}

func Deref[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

func Copy[T any](src *T) *T {
	if src == nil {
		return nil
	}
	return New(*src)
}

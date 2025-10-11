package pointer

func New[T any](value T) *T {
	return &value
}

func Deref[T any](ptr *T, def T) T {
	if ptr != nil {
		return *ptr
	}
	return def
}

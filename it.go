package iter

// ---

// New construct an It that is an iterator wrapper that enables shortcuts for composing new iterators.
func New[T any, I Iterator[T]](i I) It[T, I] {
	return It[T, I]{i}
}

// ---

// It is an iterator wrapper that enables shortcuts for composing new iterators.
type It[T any, I Iterator[T]] struct {
	it I
}

func (i It[T, I]) Next() (T, bool) {
	return i.it.Next()
}

func (i It[T, I]) Collect() []T {
	return Collect[T](i.it)
}

func (i It[T, I]) CollectInto(result []T) []T {
	return CollectInto(i.it, result)
}

// ---

var _ complete[int] = It[int, Iterator[int]]{}

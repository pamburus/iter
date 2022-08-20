package iter

// Empty returns an Iterator that has no values.
func Empty[T any]() It[T, emptyIterator[T]] {
	return New[T](emptyIterator[T]{})
}

// ---

// emptyIterator returns no values.
type emptyIterator[T any] struct{}

func (i emptyIterator[T]) It() It[T, emptyIterator[T]] {
	return New[T](i)
}

// Iter returns self.
// Next returns zero value and false.
func (emptyIterator[T]) Next() (T, bool) {
	return ZeroValue[T](), false
}

// SizeHint returns 0 and true.
func (emptyIterator[T]) SizeHint() (Size, bool) {
	return 0, true
}

// ---

type ttEmptyIt = emptyIterator[int]

var (
	_ Iterator[int] = ttEmptyIt{}
	_ SizeHinter    = ttEmptyIt{}
)

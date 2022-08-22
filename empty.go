package iter

// Empty returns an iterator that has no values.
func Empty[T any]() It[T] {
	return New[T](emptyIterator[T]{})
}

// ---

type emptyIterator[T any] struct{}

func (emptyIterator[T]) Next() (T, bool) {
	return ZeroValue[T](), false
}

func (emptyIterator[T]) SizeHint() (Size, bool) {
	return 0, true
}

// ---

type ttEmptyIt = emptyIterator[int]

var (
	_ Iterator[int] = ttEmptyIt{}
	_ SizeHinter    = ttEmptyIt{}
)

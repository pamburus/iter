package iter

// Slice returns an Iterator that iterates over elements in a slice.
func Slice[T any](values []T) It[T] {
	return New[T](&sliceIterator[T]{values})
}

// ---

// sliceIterator is an Iterator that iterates over values in a slice.
type sliceIterator[T any] struct {
	values []T
}

// Next returns next value.
func (i *sliceIterator[T]) Next() (T, bool) {
	values := i.values
	if len(values) == 0 {
		return ZeroValue[T](), false
	}

	i.values = values[1:]

	return values[0], true
}

// SizeHint returns the number of remaining values and true.
func (i *sliceIterator[T]) SizeHint() (Size, bool) {
	return Size(len(i.values)), true
}

// Collect consumes the iterator and returns the consumed values.
func (i *sliceIterator[T]) CollectAllInto(result []T) []T {
	result = append(result, (i.values)...)
	i.values = nil

	return result
}

// CollectNInto consumes the iterator and returns the consumed values.
func (i *sliceIterator[T]) CollectInto(n int, result []T) []T {
	n = MinValue(n, len(i.values))
	result = append(result, (i.values)[:n]...)
	i.values = (i.values)[n:]

	return result
}

// DiscardAll drops all elements from the iterator.
func (i *sliceIterator[T]) DiscardAll() Size {
	return i.Discard(Size(len(i.values)))
}

// Discard drops n next elements from the iterator.
func (i *sliceIterator[T]) Discard(n Size) Size {
	n = MinValue(Size(len(i.values)), n)
	i.values = (i.values)[int(n):]

	return n
}

// ---

type ttSliceIt = *sliceIterator[int]

var (
	_ Iterator[int]      = ttSliceIt(nil)
	_ SizeHinter         = ttSliceIt(nil)
	_ CollectorInto[int] = ttSliceIt(nil)
	_ Discarder          = ttSliceIt(nil)
)

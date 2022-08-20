package iter

import (
	"golang.org/x/exp/constraints"
)

// Sequence returns an Iterator that produces incrementally sequential integer numbers from begin up to to end, end is not included.
func Sequence[T constraints.Integer](begin, end T) It[T, *sequenceIterator[T]] {
	return New[T](&sequenceIterator[T]{begin, end})
}

// ---

type sequenceIterator[T constraints.Integer] struct {
	i   T
	end T
}

func (i *sequenceIterator[T]) Next() (T, bool) {
	if i.i >= i.end {
		return ZeroValue[T](), false
	}

	result := i.i
	i.i++

	return result, true
}

// SizeHint returns the number of remaining values and true.
func (i *sequenceIterator[T]) SizeHint() (Size, bool) {
	return Size(i.end - i.i), true
}

// DropAll drops all elements from the iterator.
func (i *sequenceIterator[T]) DropAll() Size {
	return i.Drop(Size(i.end - i.i))
}

// Drop drops n next elements from the iterator.
func (i *sequenceIterator[T]) Drop(n Size) Size {
	n = MinValue(Size(i.len()), n)
	i.i -= T(n)

	return n
}

// CollectInto consumes the iterator and returns the consumed values.
func (i *sequenceIterator[T]) CollectInto(result []T) []T {
	return i.CollectNInto(int(i.len()), result)
}

// CollectNInto consumes the iterator and returns the consumed values.
func (i *sequenceIterator[T]) CollectNInto(n int, result []T) []T {
	m := T(MinValue(n, int(i.len())))
	end := i.i + m
	for j := i.i; j != end; j++ {
		result = append(result, j)
	}
	i.i += m

	return result
}

func (i *sequenceIterator[T]) len() T {
	return i.end - i.i
}

// ---

type ttSequenceIt = *sequenceIterator[int]

var (
	_ Iterator[int] = ttSequenceIt(nil)
)

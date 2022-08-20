package iter

import (
	"golang.org/x/exp/constraints"
)

// Sequence returns an iterator that produces consecutive integers from begin to end, end is not included.
func Sequence[T constraints.Integer](begin, end T) It[T] {
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

func (i *sequenceIterator[T]) SizeHint() (Size, bool) {
	return Size(i.end - i.i), true
}

func (i *sequenceIterator[T]) DiscardAll() Size {
	return i.Discard(Size(i.end - i.i))
}

func (i *sequenceIterator[T]) Discard(n Size) Size {
	n = MinValue(Size(i.len()), n)
	i.i -= T(n)

	return n
}

func (i *sequenceIterator[T]) CollectAllInto(result []T) []T {
	return i.CollectInto(int(i.len()), result)
}

func (i *sequenceIterator[T]) CollectInto(n int, result []T) []T {
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
	_ SizeHinter         = ttSequenceIt(nil)
	_ Discarder          = ttSequenceIt(nil)
	_ CollectorInto[int] = ttSequenceIt(nil)
)

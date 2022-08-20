package iter

import "math"

// Take returns an iterator that limits that yields up to (but no more than) n
// values from the input iterator.
func Take[T any, I Iterator[T]](it I, n Size) It[T, *takeIterator[T, I]] {
	return New[T](&takeIterator[T, I]{it, n})
}

// ---

// takeIterator is an Iterator adapter that iterates over first N values in the underlying iterator.
type takeIterator[T any, I Iterator[T]] struct {
	it I
	n  Size
}

// Next returns next value.
func (i *takeIterator[T, I]) Next() (T, bool) {
	if i.n <= 0 {
		return ZeroValue[T](), false
	}

	result, ok := i.it.Next()
	if ok {
		i.n--
	}

	return result, ok
}

// SizeHint returns the number of remaining values and true.
func (i *takeIterator[T, I]) SizeHint() (Size, bool) {
	result := i.n
	if size, ok := SizeHint[T](i.it); ok {
		result = MinValue(result, size)
	}

	return result, true
}

// DropAll drops all elements from the iterator.
func (i *takeIterator[T, I]) DropAll() Size {
	return i.Drop(i.n)
}

// Drop drops n next elements from the iterator.
func (i *takeIterator[T, I]) Drop(n Size) Size {
	n = MinValue(n, i.n)
	m := Drop[T](i.it, n)
	i.n -= m

	return m
}

func (i *takeIterator[T, I]) CollectInto(result []T) []T {
	return i.CollectNInto(int(MinValue(i.n, Size(math.MaxInt))), result)
}

func (i *takeIterator[T, I]) CollectNInto(n int, result []T) []T {
	result = CollectNInto(i.it, int(MinValue(Size(n), i.n)), result)
	i.n -= Size(len(result))

	return result
}

// ---

type ttTakeIt = *takeIterator[int, Iterator[int]]

var (
	_ Iterator[int]       = ttTakeIt(nil)
	_ SizeHinter          = ttTakeIt(nil)
	_ Dropper             = ttTakeIt(nil)
	_ CollectorInto[int]  = ttTakeIt(nil)
	_ CollectorNInto[int] = ttTakeIt(nil)
)

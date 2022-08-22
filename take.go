package iter

import "math"

// Take returns an iterator that produces up to (but no more than) n values from the input iterator.
func Take[T any, I Iterator[T]](it I, n Size) It[T] {
	return New[T](&takeIterator[T, I]{it, n})
}

// ---

type takeIterator[T any, I Iterator[T]] struct {
	it I
	n  Size
}

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

func (i *takeIterator[T, I]) SizeHint() (Size, bool) {
	result := i.n
	if size, ok := SizeHint[T](i.it); ok {
		result = MinValue(result, size)
	}

	return result, true
}

func (i *takeIterator[T, I]) DiscardAll() Size {
	return i.Discard(i.n)
}

func (i *takeIterator[T, I]) Discard(n Size) Size {
	n = MinValue(n, i.n)
	m := Discard[T](i.it, n)
	i.n -= m

	return m
}

func (i *takeIterator[T, I]) CollectAllInto(result []T) []T {
	return i.CollectInto(int(MinValue(i.n, Size(math.MaxInt))), result)
}

func (i *takeIterator[T, I]) CollectInto(n int, result []T) []T {
	result = CollectInto(i.it, int(MinValue(Size(n), i.n)), result)
	i.n -= Size(len(result))

	return result
}

// ---

type ttTakeIt = *takeIterator[int, Iterator[int]]

var (
	_ Iterator[int]      = ttTakeIt(nil)
	_ SizeHinter         = ttTakeIt(nil)
	_ Discarder          = ttTakeIt(nil)
	_ CollectorInto[int] = ttTakeIt(nil)
)

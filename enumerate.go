package iter

import "github.com/barweiss/go-tuple"

// Enumerate returns an iterator adapter that transform each value
// into a tuple consisting of its index and the value itself.
func Enumerate[T any, I Iterator[T]](it I) It[tuple.T2[Size, T]] {
	return New[tuple.T2[Size, T]](&enumerateIterator[T, I]{it, 0})
}

// ---

type enumerateIterator[T any, I Iterator[T]] struct {
	it I
	i  Size
}

func (i *enumerateIterator[T, I]) Next() (tuple.T2[Size, T], bool) {
	if v, ok := i.it.Next(); ok {
		n := i.i
		i.i++

		return tuple.New2(n, v), true
	}

	return ZeroValue[tuple.T2[Size, T]](), false
}

func (i *enumerateIterator[T, I]) SizeHint() (Size, bool) {
	return SizeHint[T](i.it)
}

func (i *enumerateIterator[T, I]) DiscardAll() Size {
	n := DiscardAll[T](i.it)
	i.i += n

	return n
}

func (i *enumerateIterator[T, I]) Discard(n Size) Size {
	n = Discard[T](i.it, n)
	i.i += n

	return n
}

// ---

type ttEnumerateIt = *enumerateIterator[int, Iterator[int]]

var (
	_ Iterator[tuple.T2[Size, int]] = ttEnumerateIt(nil)
	_ SizeHinter                    = ttEnumerateIt(nil)
	_ Discarder                     = ttEnumerateIt(nil)
)

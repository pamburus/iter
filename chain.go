package iter

// Chain returns an Iterator adapter that Chains each value provided by the
// underlying iterator using fn.
func Chain[T any](its ...Iterator[T]) It[T, *chainIterator[T]] {
	return New[T](&chainIterator[T]{its})
}

// ---

// chainIterator is an iterator that Chains each value T provided by the underlying iterator I using function F to value R.
type chainIterator[T any] struct {
	its []Iterator[T]
}

// Next gets next value T from iterator, applies function F to it and returns returned value R.
func (i *chainIterator[T]) Next() (T, bool) {
	for !i.empty() {
		if v, ok := i.get().Next(); ok {
			return v, true
		}

		i.next()
	}

	return ZeroValue[T](), false
}

// SizeHint returns estimated remaining count of elements if known.
func (i *chainIterator[T]) SizeHint() (Size, bool) {
	result := Size(0)
	for _, it := range i.its {
		if size, ok := SizeHint[T](it); ok {
			result += size
		} else {
			return 0, false
		}
	}

	return result, true
}

// Collect consumes the iterator and returns the consumed values.
func (i *chainIterator[T]) CollectInto(result []T) []T {
	for !i.empty() {
		result = CollectInto(i.get(), result)
		i.next()
	}

	return result
}

// CollectNInto consumes the iterator and returns the consumed values.
func (i *chainIterator[T]) CollectNInto(n int, result []T) []T {
	if size, ok := i.SizeHint(); ok {
		n = int(MinValue(Size(n), size))
	}

	for !i.empty() {
		result = CollectNInto(i.get(), n-len(result), result)
		i.next()
	}

	return result
}

// DropAll drops all elements from the iterator.
func (i *chainIterator[T]) DropAll() Size {
	n := Size(0)
	for !i.empty() {
		n += DropAll[T](i.get())
		i.next()
	}

	return n
}

// Drop drops n next elements from the iterator.
func (i *chainIterator[T]) Drop(n Size) Size {
	for n > 0 && !i.empty() {
		m := Drop[T](i.get(), n)
		n -= m
		if m == 0 {
			i.next()
		}
	}

	return n
}

func (i *chainIterator[T]) empty() bool {
	return len(i.its) == 0
}

func (i *chainIterator[T]) get() Iterator[T] {
	return i.its[0]
}

func (i *chainIterator[T]) next() {
	i.its = i.its[1:]
}

// ---

type ttChainIt = *chainIterator[int]

var (
	_ Iterator[int]       = ttChainIt(nil)
	_ SizeHinter          = ttChainIt(nil)
	_ Dropper             = ttChainIt(nil)
	_ CollectorInto[int]  = ttChainIt(nil)
	_ CollectorNInto[int] = ttChainIt(nil)
)

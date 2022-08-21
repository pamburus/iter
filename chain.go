package iter

// Chain returns an iterator adapter that chains each value provided by the
// underlying iterator using fn.
func Chain[T any](its ...Iterator[T]) It[T] {
	return New[T](&chainIterator[T]{its})
}

// ---

type chainIterator[T any] struct {
	its []Iterator[T]
}

func (i *chainIterator[T]) Next() (T, bool) {
	for !i.empty() {
		if v, ok := i.get().Next(); ok {
			return v, true
		}

		i.next()
	}

	return ZeroValue[T](), false
}

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

func (i *chainIterator[T]) CollectAllInto(result []T) []T {
	for !i.empty() {
		result = CollectAllInto(i.get(), result)
		i.next()
	}

	return result
}

func (i *chainIterator[T]) CollectInto(n int, result []T) []T {
	if size, ok := i.SizeHint(); ok {
		n = int(MinValue(Size(n), size))
	}

	for !i.empty() {
		result = CollectInto(i.get(), n-len(result), result)
		i.next()
	}

	return result
}

func (i *chainIterator[T]) DiscardAll() Size {
	n := Size(0)
	for !i.empty() {
		n += DiscardAll[T](i.get())
		i.next()
	}

	return n
}

func (i *chainIterator[T]) Discard(n Size) Size {
	for n > 0 && !i.empty() {
		m := Discard[T](i.get(), n)
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
	_ Iterator[int]      = ttChainIt(nil)
	_ SizeHinter         = ttChainIt(nil)
	_ Discarder          = ttChainIt(nil)
	_ CollectorInto[int] = ttChainIt(nil)
)

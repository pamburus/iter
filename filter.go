package iter

// Filter returns an iterator adapter that filters each value provided by the
// underlying iterator using the given predicate.
func Filter[T any, I Iterator[T], P ~func(T) bool](it I, predicate P) It[T] {
	return New[T](&filterIterator[T, I, P]{it, predicate})
}

// ---

type filterIterator[T any, I Iterator[T], P ~func(T) bool] struct {
	it        I
	predicate P
}

func (i *filterIterator[T, I, P]) Next() (T, bool) {
	for v, ok := i.it.Next(); ok; v, ok = i.it.Next() {
		if i.predicate(v) {
			return v, true
		}
	}

	return ZeroValue[T](), false
}

func (i *filterIterator[T, I, P]) SizeHint() (Size, bool) {
	return SizeHint[T](i.it)
}

// ---

type ttFilterIt = *filterIterator[int, Iterator[int], func(int) bool]

var (
	_ Iterator[int] = ttFilterIt(nil)
	_ SizeHinter    = ttFilterIt(nil)
)

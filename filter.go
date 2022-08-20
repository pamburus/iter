package iter

// Filter returns an Iterator adapter that Filters each value provided by the
// underlying iterator using fn.
func Filter[T any, I Iterator[T], F ~func(T) bool](it I, fn F) It[T] {
	return New[T](&filterIterator[T, I, F]{it, fn})
}

// ---

// filterIterator is an iterator that Filters each value T provided by the underlying iterator I using function F to value R.
type filterIterator[T any, I Iterator[T], F ~func(T) bool] struct {
	it I
	fn F
}

// Next gets next value T from iterator, applies function F to it and returns returned value R.
func (i *filterIterator[T, I, F]) Next() (T, bool) {
	for v, ok := i.it.Next(); ok; v, ok = i.it.Next() {
		if i.fn(v) {
			return v, true
		}
	}

	return ZeroValue[T](), false
}

// SizeHint delegates the call to the underlying iterator.
func (i *filterIterator[T, I, F]) SizeHint() (Size, bool) {
	return SizeHint[T](i.it)
}

// ---

type ttFilterIt = *filterIterator[int, Iterator[int], func(int) bool]

var (
	_ Iterator[int] = ttFilterIt(nil)
	_ SizeHinter    = ttFilterIt(nil)
)

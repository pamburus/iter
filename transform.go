package iter

// Transform returns an iterator adapter that transforms each value provided by the underlying iterator using fn.
func Transform[T, R any, I Iterator[T], F ~func(T) R](it I, fn F) It[R] {
	return New[R](&transformIterator[T, R, I, F]{it, fn})
}

// ---

type transformIterator[T, R any, I Iterator[T], F ~func(T) R] struct {
	it I
	fn F
}

func (i transformIterator[T, R, I, F]) Next() (R, bool) {
	if v, ok := i.it.Next(); ok {
		return i.fn(v), true
	}

	return ZeroValue[R](), false
}

func (i transformIterator[T, R, I, F]) SizeHint() (Size, bool) {
	return SizeHint[T](i.it)
}

// ---

type ttTransformIt = *transformIterator[int8, int16, Iterator[int8], func(int8) int16]

var (
	_ Iterator[int16] = ttTransformIt(nil)
	_ SizeHinter      = ttTransformIt(nil)
)

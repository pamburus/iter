package iter

// Func returns a FuncIterator that calls the provided function to get next value.
func Func[T any, F ~func() (T, bool)](fn F) It[T] {
	return New[T](&funcIterator[T, F]{fn})
}

// ---

// funcIterator iterates over values of type T emitted by function F.
type funcIterator[T any, F ~func() (T, bool)] struct {
	fn F
}

// Next returns next value.
func (i funcIterator[T, F]) Next() (T, bool) {
	return i.fn()
}

// ---

type ttFuncIt = *funcIterator[int, func() (int, bool)]

var (
	_ Iterator[int] = ttFuncIt(nil)
)

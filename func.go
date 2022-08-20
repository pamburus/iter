package iter

// Func returns an iterator that calls the provided function to get next value.
func Func[T any, F ~func() (T, bool)](fn F) It[T] {
	return New[T](&funcIterator[T, F]{fn})
}

// ---

type funcIterator[T any, F ~func() (T, bool)] struct {
	fn F
}

func (i funcIterator[T, F]) Next() (T, bool) {
	return i.fn()
}

// ---

type ttFuncIt = *funcIterator[int, func() (int, bool)]

var (
	_ Iterator[int] = ttFuncIt(nil)
)

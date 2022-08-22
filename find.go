package iter

// Find searches for an element of an iterator that satisfies a predicate.
//
// It takes a predicate that returns true or false and applies it to each element of the iterator,
// and if any of them return true, then Find returns the element and true.
// If they all return false, it returns zero value and false.
//
// It will stop processing as soon as the predicate returns true.
func Find[T any, I Iterator[T], P ~func(T) bool](it I, predicate P) (T, bool) {
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		if predicate(v) {
			return v, true
		}
	}

	return ZeroValue[T](), false
}

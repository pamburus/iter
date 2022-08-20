package iter

// Any tests if any element of the iterator matches a predicate.
//
// It takes a predicate that returns true or false.
// It applies this closure to each element of the iterator, and if any of them return true, then so does Any.
// If they all return false, it returns false.
//
// It will stop processing as soon as it finds a true, given that no matter what else happens, the result will also be true.
//
// An empty iterator returns false.
func Any[T any, I Iterator[T], P ~func(T) bool](it I, predicate P) bool {
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		if predicate(v) {
			return true
		}
	}

	return false
}

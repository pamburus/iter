package iter

// All tests if every element of the iterator matches a predicate.
//
// It takes a predicate that returns true or false.
// It applies this predicate to each element of the iterator, and if they all return true, then so does All.
// If any of them return false, it returns false.
//
// It will stop processing as soon as it finds a false, given that no matter what else happens, the result will also be false.
//
// An empty iterator returns true.
func All[T any, I Iterator[T], P ~func(T) bool](it I, predicate P) bool {
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		if !predicate(v) {
			return false
		}
	}

	return true
}

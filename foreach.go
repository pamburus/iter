package iter

// ForEach consumes the iterator applying fn to each produced value.
func ForEach[T any, I Iterator[T], F ~func(T)](it I, fn F) {
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		fn(v)
	}
}

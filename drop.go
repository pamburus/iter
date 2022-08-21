package iter

// Discard consumes up to n elements from an iterator returning the number of consumed elements.
func Discard[T any, I Iterator[T]](it I, n Size) Size {
	if discarder, ok := Iterator[T](it).(Discarder); ok {
		return discarder.Discard(n)
	}

	i := Size(0)
	ForEach(Take[T](it, n), func(T) {
		i++
	})

	return i
}

// DiscardAll consumes all elements from an iterator returning the number of consumed elements.
func DiscardAll[T any, I Iterator[T]](it I) Size {
	if discarder, ok := Iterator[T](it).(Discarder); ok {
		return discarder.DiscardAll()
	}

	i := Size(0)
	ForEach(it, func(T) {
		i++
	})

	return i
}

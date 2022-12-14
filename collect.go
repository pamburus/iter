package iter

// CollectAll transforms an iterator into a slice.
func CollectAll[T any, I Iterator[T]](it I) []T {
	n := 0
	if size, ok := SizeHint[T](it); ok {
		n = int(size)
	}

	result := make([]T, 0, n)

	return CollectAllInto(it, result)
}

// CollectAllInto gets all values from an iterator and appends them to a given slice returning the same or a new slice.
func CollectAllInto[T any, I Iterator[T], S ~[]T](it I, result S) S {
	if coll, ok := Iterator[T](it).(CollectorInto[T]); ok {
		return coll.CollectAllInto(result)
	}

	ForEach(it, func(v T) {
		result = append(result, v)
	})

	return result
}

// Collect transforms up to n values of an iterator into a slice.
func Collect[T any, I Iterator[T]](it I, n int) []T {
	if size, ok := SizeHint[T](it); ok {
		n = int(MinValue(Size(n), size))
	}

	result := make([]T, 0, n)

	return CollectInto(it, n, result)
}

// CollectInto gets up to N values from an iterator and appends them to a given slice returning the same or a new slice.
func CollectInto[T any, I Iterator[T], S ~[]T](it I, n int, result S) S {
	if coll, ok := Iterator[T](it).(CollectorInto[T]); ok {
		return coll.CollectInto(n, result)
	}

	for n > 0 {
		if v, ok := it.Next(); ok {
			result = append(result, v)
		} else {
			break
		}
		n--
	}

	return result
}

package iter

// Collect transforms an Iterator into a slice.
func Collect[T any, I Iterator[T]](it I) []T {
	n := 0
	if size, ok := SizeHint[T](it); ok {
		n = int(size)
	}

	result := make([]T, 0, n)

	return CollectInto(it, result)
}

// CollectInto gets all values from an Iterator and appends them to a given slice returning a new slice.
func CollectInto[T any, I Iterator[T], S ~[]T](it I, result S) S {
	if coll, ok := Iterator[T](it).(CollectorInto[T]); ok {
		return coll.CollectInto(result)
	}

	ForEach(it, func(v T) {
		result = append(result, v)
	})

	return result
}

// CollectN transforms up to n values of an Iterator into a slice.
func CollectN[T any, I Iterator[T]](it I, n int) []T {
	if size, ok := SizeHint[T](it); ok {
		n = int(MinValue(Size(n), size))
	}

	result := make([]T, 0, n)

	return CollectNInto(it, n, result)
}

// CollectNInto gets up to N values from an Iterator and appends them to a given slice returning a new slice.
func CollectNInto[T any, I Iterator[T], S ~[]T](it I, n int, result S) S {
	if coll, ok := Iterator[T](it).(CollectorNInto[T]); ok {
		return coll.CollectNInto(n, result)
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

package iter

// Collect transforms an Iterator into a slice.
func Collect[T any, I Iterator[T]](it I) []T {
	result := make([]T, 0)

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

package iter

// Drop consumes up to n elements from an iterator returning the number of consumed elements.
func Drop[T any, I Iterator[T]](it I, n Size) Size {
	if dropper, ok := Iterator[T](it).(Dropper); ok {
		return dropper.Drop(n)
	}

	i := Size(0)
	for ; i < n; i++ {
		if _, ok := it.Next(); ok {
			i++
		} else {
			break
		}
	}

	return i
}

// DropAll consumes all elements from an iterator returning the number of consumed elements.
func DropAll[T any, I Iterator[T]](it I) Size {
	if dropper, ok := Iterator[T](it).(Dropper); ok {
		return dropper.DropAll()
	}

	i := Size(0)
	ForEach(it, func(T) {
		i++
	})

	return i
}

package iter

// ---

// SizeHint returns estimated number of items left in the iterator if known.
func SizeHint[T any, I Iterator[T]](it I) (Size, bool) {
	if hinter, ok := Iterator[T](it).(SizeHinter); ok {
		return hinter.SizeHint()
	}

	return 0, false
}

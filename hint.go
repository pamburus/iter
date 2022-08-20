package iter

// ---

// SizeHint consumes an Iterator returning count of obtained values.
func SizeHint[T any, I Iterator[T]](it I) (Size, bool) {
	if hinter, ok := Iterator[T](it).(SizeHinter); ok {
		return hinter.SizeHint()
	}

	return 0, false
}

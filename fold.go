package iter

// Fold takes initial value as the result value and then applies function F
// to the result value and and next value from Iterator and updates the result value.
func Fold[T, R any, I Iterator[T], F ~func(R, T) R](it I, initial R, f F) R {
	result := initial
	ForEach(it, func(v T) {
		result = f(result, v)
	})

	return result
}

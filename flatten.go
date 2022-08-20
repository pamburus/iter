package iter

// Flatten returns an Iterator adapter that Flattens each value provided by the
// underlying iterator using fn.
func Flatten[T any, I Iterator[Iterator[T]]](it I) It[T] {
	return New[T](&flattenIterator[T, I]{it, nil})
}

// ---

// flattenIterator is an iterator that Flattens each value T provided by the underlying iterator I using function F to value R.
type flattenIterator[T any, I Iterator[Iterator[T]]] struct {
	it  I
	cur Iterator[T]
}

// Next gets next value T from iterator, applies function F to it and returns returned value R.
func (i *flattenIterator[T, I]) Next() (T, bool) {
	for {
		if i.cur == nil {
			if v, ok := i.it.Next(); ok {
				i.cur = v
			} else {
				break
			}
		}

		if i.cur != nil {
			if v, ok := (i.cur).Next(); ok {
				return v, true
			}
			i.cur = nil
		}
	}

	return ZeroValue[T](), false
}

// ---

type ttFlattenIt = *flattenIterator[int, Iterator[Iterator[int]]]

var (
	_ Iterator[int] = ttFlattenIt(nil)
)

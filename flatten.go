package iter

// Flatten returns an iterator of values made from iterator of iterators of values.
func Flatten[T any, I Iterator[Iterator[T]]](it I) It[T] {
	return New[T](&flattenIterator[T, I]{it, nil})
}

// ---

type flattenIterator[T any, I Iterator[Iterator[T]]] struct {
	it  I
	cur Iterator[T]
}

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

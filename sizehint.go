package iter

// WithSizeHint returns a wrapped Iterator that overrides SizeHint.
func WithSizeHint[T any, I Iterator[T]](it I, n Size) It[T, *sizeHintIterator[T, I]] {
	return New[T](&sizeHintIterator[T, I]{New[T](it), n})
}

// ---

type sizeHintIterator[T any, I Iterator[T]] struct {
	It[T, I]
	n Size
}

func (i *sizeHintIterator[T, F]) Iter() Iterator[T] {
	return i
}

func (i *sizeHintIterator[T, F]) SizeHint() (Size, bool) {
	return i.n, true
}

// ---

type ttSizeHintIt = *sizeHintIterator[int, Iterator[int]]

var (
	_ Iterator[int]       = ttSizeHintIt(nil)
	_ SizeHinter          = ttSizeHintIt(nil)
	_ Dropper             = ttSizeHintIt(nil)
	_ CollectorInto[int]  = ttSizeHintIt(nil)
	_ CollectorNInto[int] = ttSizeHintIt(nil)
)

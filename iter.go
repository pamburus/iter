package iter

// Size is a number of elements in iterator or an index of element in iterator.
type Size = uint64

// Iterator allows iteration over an abstract stream of items.
type Iterator[T any] interface {
	Next() (T, bool)
}

// CollectorInto is any type the can have CollectInto method taking a slice of T and returning a slice of T.
type CollectorInto[T any] interface {
	CollectInto([]T) []T
}

// ---

type complete[T any] interface {
	Iterator[T]
	CollectorInto[T]
}

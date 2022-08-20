package iter

// Size is a number of items in an iterator or an index of item in iterator.
type Size = uint64

// Iterator allows iteration over an abstract stream of items.
type Iterator[T any] interface {
	Next() (T, bool)
}

// Iterable allows getting an iterator over some items.
type Iterable[T any] interface {
	Iter() Iterator[T]
}

// SizeHinter allows getting estimated remaining number of items in an iterator.
type SizeHinter interface {
	SizeHint() (Size, bool)
}

// CollectorInto allows collecting items from an iterator in a slice.
type CollectorInto[T any] interface {
	CollectInto(int, []T) []T
	CollectAllInto([]T) []T
}

// Discarder allows dropping items from an iterator.
type Discarder interface {
	Discard(Size) Size
	DiscardAll() Size
}

// ---

type complete[T any] interface {
	Iterator[T]
	Iterable[T]
	SizeHinter
	Discarder
	CollectorInto[T]
}

package iter

// Size is a number of elements in iterator or an index of element in iterator.
type Size = uint64

// Iterator allows iteration over an abstract stream of items.
type Iterator[T any] interface {
	Next() (T, bool)
}

// Iterable allows getting an Iterator.
type Iterable[T any] interface {
	Iter() Iterator[T]
}

// SizeHinter allows to optionally get estimated remaining number of items in an Iterator.
type SizeHinter interface {
	SizeHint() (Size, bool)
}

// CollectorInto is any type the can have CollectInto method taking a slice of T and returning a slice of T.
type CollectorInto[T any] interface {
	CollectInto([]T) []T
}

// CollectorNInto is any type the can have CollectInto method taking a slice of T and returning a slice of T.
type CollectorNInto[T any] interface {
	CollectNInto(int, []T) []T
}

// Dropper allows dropping elements from iterator.
type Dropper interface {
	Drop(Size) Size
	DropAll() Size
}

// ---

type complete[T any] interface {
	Iterator[T]
	Iterable[T]
	SizeHinter
	Dropper
	CollectorInto[T]
	CollectorNInto[T]
}

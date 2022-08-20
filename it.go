package iter

import "github.com/pamburus/iter/optional"

// ---

// New construct an It that is an iterator wrapper that enables shortcuts for composing new iterators.
func New[T any](i Iterator[T]) It[T] {
	return It[T]{i}
}

// ---

// It is an iterator wrapper that enables shortcuts for composing new iterators.
type It[T any] struct {
	it Iterator[T]
}

func (i It[T]) Iter() Iterator[T] {
	return i.it
}

func (i It[T]) Next() (T, bool) {
	return i.it.Next()
}

func (i It[T]) All(predicate func(T) bool) bool {
	return All(i.it, predicate)
}

func (i It[T]) Any(predicate func(T) bool) bool {
	return Any(i.it, predicate)
}

func (i It[T]) Find(predicate func(T) bool) (T, bool) {
	return Find(i.it, predicate)
}

func (i It[T]) Collect() []T {
	return Collect[T](i.it)
}

func (i It[T]) CollectInto(result []T) []T {
	return CollectInto(i.it, result)
}

func (i It[T]) CollectN(n int) []T {
	return CollectN[T](i.it, n)
}

func (i It[T]) CollectNInto(n int, result []T) []T {
	return CollectNInto(i.it, n, result)
}

func (i It[T]) Take(n Size) It[T] {
	return Take[T](i.it, n)
}

func (i It[T]) Drop(n Size) Size {
	return Drop[T](i.it, n)
}

func (i It[T]) DropAll() Size {
	return DropAll[T](i.it)
}

func (i It[T]) Filter(predicate func(T) bool) It[T] {
	return Filter(i.it, predicate)
}

func (i It[T]) Chain(other ...Iterator[T]) It[T] {
	return Chain(append([]Iterator[T]{i.it}, other...)...)
}

func (i It[T]) Reduce(f func(T, T) T) optional.Value[T] {
	return Reduce(i.it, f)
}

func (i It[T]) MinBy(less func(T, T) bool) optional.Value[T] {
	return MinBy(i.it, less)
}

func (i It[T]) MaxBy(less func(T, T) bool) optional.Value[T] {
	return MaxBy(i.it, less)
}

func (i It[T]) SizeHint() (Size, bool) {
	return SizeHint[T](i.it)
}

// ---

var _ complete[int] = It[int]{}

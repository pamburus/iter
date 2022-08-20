package iter

import "github.com/pamburus/iter/optional"

// ---

// New construct an It that is an iterator wrapper that enables shortcuts for composing new iterators.
func New[T any, I Iterator[T]](i I) It[T, I] {
	return It[T, I]{i}
}

// ---

// It is an iterator wrapper that enables shortcuts for composing new iterators.
type It[T any, I Iterator[T]] struct {
	it I
}

func (i It[T, I]) It() It[T, Iterator[T]] {
	return New[T](Iterator[T](i.it))
}

func (i It[T, I]) Iter() Iterator[T] {
	return i.it
}

func (i It[T, I]) Next() (T, bool) {
	return i.it.Next()
}

func (i It[T, I]) All(predicate func(T) bool) bool {
	return All(i.it, predicate)
}

func (i It[T, I]) Any(predicate func(T) bool) bool {
	return Any(i.it, predicate)
}

func (i It[T, I]) Find(predicate func(T) bool) (T, bool) {
	return Find(i.it, predicate)
}

func (i It[T, I]) Collect() []T {
	return Collect[T](i.it)
}

func (i It[T, I]) CollectInto(result []T) []T {
	return CollectInto(i.it, result)
}

func (i It[T, I]) CollectN(n int) []T {
	return CollectN[T](Iterator[T](i.it), n)
}

func (i It[T, I]) CollectNInto(n int, result []T) []T {
	return CollectNInto(Iterator[T](i.it), n, result)
}

func (i It[T, I]) Take(n Size) It[T, *takeIterator[T, Iterator[T]]] {
	return Take[T](Iterator[T](i.it), n)
}

func (i It[T, I]) Drop(n Size) Size {
	return Drop[T](Iterator[T](i.it), n)
}

func (i It[T, I]) DropAll() Size {
	return DropAll[T](Iterator[T](i.it))
}

func (i It[T, I]) Filter(predicate func(T) bool) It[T, *filterIterator[T, Iterator[T], func(T) bool]] {
	return Filter(Iterator[T](i.it), predicate)
}

func (i It[T, I]) Chain(other ...Iterator[T]) It[T, *chainIterator[T]] {
	return Chain(append([]Iterator[T]{i.it}, other...)...)
}

func (i It[T, I]) Reduce(f func(T, T) T) optional.Value[T] {
	return Reduce(i.it, f)
}

func (i It[T, I]) MinBy(less func(T, T) bool) optional.Value[T] {
	return MinBy(i.it, less)
}

func (i It[T, I]) MaxBy(less func(T, T) bool) optional.Value[T] {
	return MaxBy(i.it, less)
}

func (i It[T, I]) SizeHint() (Size, bool) {
	return SizeHint[T](i.it)
}

func (i It[T, I]) WithSizeHint(n Size) It[T, *sizeHintIterator[T, Iterator[T]]] {
	return WithSizeHint[T](Iterator[T](i.it), n)
}

// ---

var _ complete[int] = It[int, Iterator[int]]{}

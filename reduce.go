package iter

import (
	"github.com/pamburus/iter/optional"
	"golang.org/x/exp/constraints"
)

// Reduce reduces iterator using provided function f.
func Reduce[T any, I Iterator[T], F ~func(T, T) T](it I, f F) optional.Value[T] {
	first, ok := it.Next()
	if !ok {
		return optional.None[T]()
	}

	return optional.Some(Fold(it, first, f))
}

// Min returns minimum value from the iterator if there are any.
func Min[T constraints.Ordered, I Iterator[T]](it I) optional.Value[T] {
	return Reduce(it, MinValue[T])
}

// Max returns maximum value from the iterator if there are any.
func Max[T constraints.Ordered, I Iterator[T]](it I) optional.Value[T] {
	return Reduce(it, MaxValue[T])
}

// MinBy returns minimum value from the iterator if there are any.
// Comparison of two values is performed by the provided less function.
func MinBy[T any, I Iterator[T], L ~func(T, T) bool](it I, less L) optional.Value[T] {
	return Reduce(it, UsingLess(less).Min())
}

// MaxBy returns maximum value from the iterator if there are any.
// Comparison of two values is performed by the provided less function.
func MaxBy[T any, I Iterator[T], L ~func(T, T) bool](it I, less L) optional.Value[T] {
	return Reduce(it, UsingLess(less).Max())
}

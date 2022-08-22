package iter

import "golang.org/x/exp/constraints"

// ---

// ZeroValue returns a zero-initialized value of type T.
func ZeroValue[T any]() T {
	var result T

	return result
}

// IsZeroValue returns true if the value is zero initialized.
func IsZeroValue[T comparable](value T) bool {
	return value == ZeroValue[T]()
}

// ---

// MinValue returns minimum value of two values.
func MinValue[T constraints.Ordered](left, right T) T {
	if left < right {
		return left
	}

	return right
}

// MaxValue returns maximum value of two values.
func MaxValue[T constraints.Ordered](left, right T) T {
	if left > right {
		return left
	}

	return right
}

// ---

// UsingLess returns a build of comparison functions based on the provided less function.
func UsingLess[T any, L ~func(T, T) bool](less L) LessBuilder[T, L] {
	return LessBuilder[T, L]{less}
}

// LessBuilder allows to compose various comparison functions based on the stored less function.
type LessBuilder[T any, L ~func(T, T) bool] struct {
	less L
}

// Min returns a function that will return a minimum of two values comparing them by the stored less function.
func (b LessBuilder[T, L]) Min() func(T, T) T {
	return func(left, right T) T {
		if b.less(left, right) {
			return left
		}

		return right
	}
}

// Max returns a function that will return a maximum of two values comparing them by the stored less function.
func (b LessBuilder[T, L]) Max() func(T, T) T {
	return func(left, right T) T {
		if b.less(right, left) {
			return left
		}

		return right
	}
}

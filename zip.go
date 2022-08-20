package iter

import (
	"github.com/barweiss/go-tuple"
	"github.com/pamburus/iter/optional"
)

// Zip returns an Iterator adapter that Zips each value provided by the
// underlying iterator using fn.
func Zip[T, U any, I1 Iterator[T], I2 Iterator[U]](i1 I1, i2 I2) It[tuple.T2[T, U]] {
	return New[tuple.T2[T, U]](
		&zipIterator[T, U, I1, I2]{
			i1,
			i2,
			optional.None[T](),
			optional.None[U](),
		},
	)
}

// ---

// zipIterator is an iterator that Zips each value T provided by the underlying iterator I using function F to value R.
type zipIterator[T, U any, I1 Iterator[T], I2 Iterator[U]] struct {
	it1 I1
	it2 I2
	v1  optional.Value[T]
	v2  optional.Value[U]
}

// Next gets next value T from iterator, applies function F to it and returns returned value R.
func (i *zipIterator[T, U, I1, I2]) Next() (tuple.T2[T, U], bool) {
	if i.v1.IsNone() {
		i.v1 = optional.New(i.it1.Next())
	}
	if v1, ok := i.v1.Unwrap(); ok {
		if i.v2.IsNone() {
			i.v2 = optional.New(i.it2.Next())
		}
		if v2, ok := i.v2.Unwrap(); ok {
			i.v1.Reset()
			i.v2.Reset()

			return tuple.New2(v1, v2), true
		}
	}

	return ZeroValue[tuple.T2[T, U]](), false
}

func (i *zipIterator[T, U, I1, I2]) SizeHint() (Size, bool) {
	if s1, ok := SizeHint[T](i.it1); ok {
		if s2, ok := SizeHint[U](i.it2); ok {
			return MinValue(s1, s2), true
		}
	}

	return 0, false
}

// ---

type ttZipIt = *zipIterator[int, uint, Iterator[int], Iterator[uint]]

var (
	_ Iterator[tuple.T2[int, uint]] = ttZipIt(nil)
	_ SizeHinter                    = ttZipIt(nil)
)

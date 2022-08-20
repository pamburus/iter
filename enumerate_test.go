package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestEnumerate(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(
		iter.Enumerate[int](
			iter.Sequence(10, 13),
		).Collect(),
	).To(
		Equal([]tuple.T2[uint64, int]{
			{V1: 0, V2: 10},
			{V1: 1, V2: 11},
			{V1: 2, V2: 12},
		}),
	)
}

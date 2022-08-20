package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestZip(t *testing.T) {
	g := gomegax.New(t)

	g.Run("EqualSize", func(g gomegax.G) {
		g.Expect(
			iter.Zip[int, int](
				iter.Sequence(0, 3),
				iter.Sequence(10, 13),
			).CollectAll(),
		).To(
			Equal([]tuple.T2[int, int]{
				{V1: 0, V2: 10},
				{V1: 1, V2: 11},
				{V1: 2, V2: 12},
			}),
		)
	})

	g.Run("LeftSizeLess", func(g gomegax.G) {
		g.Expect(
			iter.Zip[int, int](
				iter.Sequence(0, 3),
				iter.Sequence(10, 15),
			).CollectAll(),
		).To(
			Equal([]tuple.T2[int, int]{
				{V1: 0, V2: 10},
				{V1: 1, V2: 11},
				{V1: 2, V2: 12},
			}),
		)
	})

	g.Run("LeftSizeMore", func(g gomegax.G) {
		g.Expect(
			iter.Zip[int, int](
				iter.Sequence(0, 5),
				iter.Sequence(10, 13),
			).CollectAll(),
		).To(
			Equal([]tuple.T2[int, int]{
				{V1: 0, V2: 10},
				{V1: 1, V2: 11},
				{V1: 2, V2: 12},
			}),
		)
	})
}

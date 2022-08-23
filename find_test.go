package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
	"github.com/pamburus/optional"
)

func TestFind(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(
		optional.New(iter.Sequence(10, 13).Find(func(value int) bool {
			return value > 10
		})),
	).To(
		Equal(optional.Some(11)),
	)

	g.Expect(
		optional.New(iter.Sequence(10, 13).Find(func(value int) bool {
			return value > 12
		})),
	).To(
		Equal(optional.None[int]()),
	)
}

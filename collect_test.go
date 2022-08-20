package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestCollect(t *testing.T) {
	g := gomegax.New(t)

	g.Run("Empty", func(g gomegax.G) {
		g.Expect(iter.Collect[int](iter.Slice[int](nil))).To(Equal([]int{}))
	})
}

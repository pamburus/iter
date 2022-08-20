package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
	"github.com/pamburus/iter/optional"
)

func TestFunc(t *testing.T) {
	g := gomegax.New(t)

	f := func() func() (int, bool) {
		i := 0

		return func() (int, bool) {
			if i >= 3 {
				return 0, false
			}

			i++

			return i, true
		}
	}

	g.Run("Raw", func(g gomegax.G) {
		it := iter.Func(f())
		g.Expect(optional.New(it.Next())).To(Equal(optional.New(1, true)))
		g.Expect(optional.New(it.Next())).To(Equal(optional.New(2, true)))
		g.Expect(optional.New(it.Next())).To(Equal(optional.New(3, true)))
		g.Expect(optional.New(it.Next())).To(Equal(optional.New(0, false)))
		g.Expect(optional.New(it.Next())).To(Equal(optional.New(0, false)))
	})

	g.Run("Collect", func(g gomegax.G) {
		g.Expect(iter.Collect[int](iter.Func(f()))).To(Equal([]int{1, 2, 3}))
		g.Expect(iter.Func(f()).Collect()).To(Equal([]int{1, 2, 3}))
	})

	g.Run("CollectN", func(g gomegax.G) {
		g.Expect(iter.Func(f()).CollectN(2)).To(Equal([]int{1, 2}))
	})
}
package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestReduce(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(iter.Max[int](iter.Slice([]int{16, 2, 43, 42, 25})).Value()).To(Equal(43))
}

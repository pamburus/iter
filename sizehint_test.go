package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestSizeHint(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(iter.WithSizeHint[int](iter.Sequence(1, 6), 3).Collect()).To(Equal([]int{1, 2, 3, 4, 5}))
}

func BenchmarkSizeHintA(b *testing.B) {
	for i := 0; i != b.N; i++ {
		iter.DropAll[int](iter.WithSizeHint[int](iter.Sequence(1, 10000), 3))
	}
}

func BenchmarkSizeHintB(b *testing.B) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(1, 10000).WithSizeHint(3).DropAll()
	}
}

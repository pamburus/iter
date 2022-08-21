package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
	"github.com/pamburus/iter/optional"
)

func TestSequence(t *testing.T) {
	g := gomegax.New(t)

	it := iter.Sequence(1, 4)
	g.Expect(optional.New(it.Next())).To(Equal(optional.New(1, true)))
	g.Expect(optional.New(it.Next())).To(Equal(optional.New(2, true)))
	g.Expect(optional.New(it.Next())).To(Equal(optional.New(3, true)))
	g.Expect(optional.New(it.Next())).To(Equal(optional.New(0, false)))
	g.Expect(optional.New(it.Next())).To(Equal(optional.New(0, false)))
}

func BenchmarkSequence(b *testing.B) {
	funcs := []tuple.T2[string, func(*testing.B, int)]{
		tuple.New2("D:v1", benchmarkSequenceDiscardA),
		tuple.New2("D:v2", benchmarkSequenceDiscardB),
		tuple.New2("C:v1", benchmarkSequenceCollectA),
		tuple.New2("C:v2", benchmarkSequenceCollectB),
	}

	sizes := []tuple.T2[string, int]{
		tuple.New2("1e1", 10),
		tuple.New2("1e4", 10000),
	}

	for _, n := range sizes {
		b.Run(n.V1, func(b *testing.B) {
			for _, fn := range funcs {
				b.Run(fn.V1, func(b *testing.B) {
					fn.V2(b, n.V2)
				})
			}
		})
	}
}

func benchmarkSequenceDiscardA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.DiscardAll[int](iter.Sequence(0, n))
	}
}

func benchmarkSequenceDiscardB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n).DiscardAll()
	}
}

func benchmarkSequenceCollectA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.CollectAll[int](iter.Sequence(0, n))
	}
}

func benchmarkSequenceCollectB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n).CollectAll()
	}
}

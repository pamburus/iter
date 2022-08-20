package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestChain(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(iter.Sequence(1, 4).Chain(iter.Sequence(10, 12)).CollectAll()).To(Equal([]int{1, 2, 3, 10, 11}))
}

func BenchmarkChain(b *testing.B) {
	funcs := []tuple.T2[string, func(*testing.B, int)]{
		tuple.New2("D:v1", benchmarkChainDropA),
		tuple.New2("D:v2", benchmarkChainDropB),
		tuple.New2("C:v1", benchmarkChainCollectA),
		tuple.New2("C:v2", benchmarkChainCollectB),
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

func benchmarkChainDropA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.DropAll[int](iter.Chain[int](iter.Sequence(0, n/2), iter.Sequence(n/2, n)))
	}
}

func benchmarkChainDropB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n/2).Chain(iter.Sequence(n/2, n)).DropAll()
	}
}

func benchmarkChainCollectA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.CollectAll[int](iter.Chain[int](iter.Sequence(0, n/2), iter.Sequence(n/2, n)))
	}
}

func benchmarkChainCollectB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n/2).Chain(iter.Sequence(n/2, n)).CollectAll()
	}
}

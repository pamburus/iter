package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	"github.com/pamburus/iter"
)

func BenchmarkSlice(b *testing.B) {
	funcs := []tuple.T2[string, func(*testing.B, int)]{
		tuple.New2("D:v1", benchmarkSliceDiscardA),
		tuple.New2("D:v2", benchmarkSliceDiscardB),
		tuple.New2("C:v1", benchmarkSliceCollectA),
		tuple.New2("C:v2", benchmarkSliceCollectB),
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

func benchmarkSliceDiscardA(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.DiscardAll[int](iter.Slice(data))
	}
}

func benchmarkSliceDiscardB(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.Slice(data).DiscardAll()
	}
}

func benchmarkSliceCollectA(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.CollectAll[int](iter.Slice(data))
	}
}

func benchmarkSliceCollectB(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.Slice(data).CollectAll()
	}
}

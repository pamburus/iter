package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	"github.com/pamburus/iter"
)

func BenchmarkFilter(b *testing.B) {
	funcs := []tuple.T2[string, func(*testing.B, int)]{
		tuple.New2("D:v1", benchmarkFilterDiscardA),
		tuple.New2("D:v2", benchmarkFilterDiscardB),
		tuple.New2("D:v3", benchmarkFilterDiscardC),
		tuple.New2("C:v1", benchmarkFilterCollectA),
		tuple.New2("C:v2", benchmarkFilterCollectB),
		tuple.New2("C:v3", benchmarkFilterCollectC),
		tuple.New2("gold", benchmarkFilterCollectIdeal),
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

func benchmarkFilterDiscardA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.DiscardAll[int](iter.Filter(iter.Sequence(0, n), odd))
	}
}

func benchmarkFilterDiscardB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n).Filter(odd).DiscardAll()
	}
}

func benchmarkFilterDiscardC(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.Slice(data).Filter(odd).DiscardAll()
	}
}

func benchmarkFilterCollectA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.CollectAll[int](iter.Filter(iter.Sequence(0, n), odd))
	}
}

func benchmarkFilterCollectB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n).Filter(odd).CollectAll()
	}
}

func benchmarkFilterCollectC(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.Slice(data).Filter(odd).CollectAll()
	}
}

func benchmarkFilterCollectIdeal(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		buf := make([]int, 0)
		for j := 0; j != n; j++ {
			if odd(j) {
				buf = append(buf, j)
			}
		}
		_ = buf
	}
}

func odd(value int) bool {
	return value%2 != 0
}

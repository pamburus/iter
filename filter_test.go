package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	"github.com/pamburus/iter"
)

func BenchmarkFilter(b *testing.B) {
	funcs := []tuple.T2[string, func(*testing.B, int)]{
		tuple.New2("D:v1", benchmarkFilterDropA),
		tuple.New2("D:v2", benchmarkFilterDropB),
		tuple.New2("D:v3", benchmarkFilterDropC),
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

func benchmarkFilterDropA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.DropAll[int](iter.Filter(iter.Sequence(0, n), odd))
	}
}

func benchmarkFilterDropB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Sequence(0, n).Filter(odd).DropAll()
	}
}

func benchmarkFilterDropC(b *testing.B, n int) {
	data := iter.Sequence(0, n).CollectAll()
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		iter.Slice(data).Filter(odd).DropAll()
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

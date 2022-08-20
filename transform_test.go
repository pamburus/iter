package iter_test

import (
	"testing"

	tuple "github.com/barweiss/go-tuple"
	"github.com/pamburus/iter"
)

func BenchmarkTransform(b *testing.B) {
	funcs := []tuple.T2[string, func(*testing.B, int)]{
		tuple.New2("D:v1", benchmarkTransformDropA),
		tuple.New2("D:v2", benchmarkTransformDropB),
		tuple.New2("C:v1", benchmarkTransformCollectA),
		tuple.New2("C:v2", benchmarkTransformCollectB),
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

func benchmarkTransformDropA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.DropAll[int](iter.Transform(iter.Sequence(0, n), double))
	}
}

func benchmarkTransformDropB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Transform(iter.Sequence(0, n), double).DropAll()
	}
}

func benchmarkTransformCollectA(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Collect[int](iter.Transform(iter.Sequence(0, n), double))
	}
}

func benchmarkTransformCollectB(b *testing.B, n int) {
	for i := 0; i != b.N; i++ {
		iter.Transform(iter.Sequence(0, n), double).Collect()
	}
}

func double(value int) int {
	return value * 2
}

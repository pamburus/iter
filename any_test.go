package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestAny(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(
		iter.Sequence(10, 13).Any(func(value int) bool {
			return value > 11
		}),
	).To(BeTrue())

	g.Expect(
		iter.Sequence(10, 13).Any(func(value int) bool {
			return value > 13
		}),
	).To(BeFalse())
}

package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestAll(t *testing.T) {
	g := gomegax.New(t)

	g.Expect(
		iter.Sequence(10, 13).All(func(value int) bool {
			return value >= 10
		}),
	).To(BeTrue())

	g.Expect(
		iter.Sequence(10, 13).All(func(value int) bool {
			return value > 10
		}),
	).To(BeFalse())
}

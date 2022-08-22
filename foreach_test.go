package iter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pamburus/gomegax"
	"github.com/pamburus/iter"
)

func TestForEach(t *testing.T) {
	g := gomegax.New(t)

	type testCase struct {
		name   string
		source []int
	}

	testCases := []testCase{
		{"Empty", []int{}},
		{"1-Item", []int{42}},
		{"2-Items", []int{42, 43}},
		{"3-Items", []int{43, 42, 43}},
	}

	for _, testCase := range testCases {
		g.Run(testCase.name, func(g gomegax.G) {
			result := []int{}
			iter.ForEach(iter.Slice(testCase.source), func(i int) {
				result = append(result, i)
			})
			g.Expect(result).To(Equal(testCase.source))
		})
	}
}

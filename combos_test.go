package combos_test

import (
	"testing"

	"github.com/notnil/combos"
)

type combinationTest struct {
	n        int
	r        int
	expected int
}

var combinationTests = []combinationTest{
	{7, 5, 21},
	{7, 4, 35},
	{10, 2, 45},
	{100, 2, 4950},
	{1000, 2, 499500},
}

func TestCombinations(t *testing.T) {
	for _, ct := range combinationTests {
		actual := combos.New(ct.n, ct.r)
		if len(actual) != ct.expected {
			t.Errorf("%dC%d: expected %d, actual %d", ct.n, ct.r, ct.expected, len(actual))
		}
	}
}

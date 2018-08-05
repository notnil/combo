package combos_test

import (
	"testing"

	"github.com/notnil/combos"
)

func Benchmark10C2(b *testing.B) { benchmarknCr(10, 2, b) }

func Benchmark100C2(b *testing.B) { benchmarknCr(100, 2, b) }

func Benchmark1000C2(b *testing.B) { benchmarknCr(1000, 2, b) }

func benchmarknCr(n int, r int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		combos.New(n, r)
	}
}

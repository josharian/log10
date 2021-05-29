package log10_test

import (
	"fmt"
	"math"
	"math/bits"
	"testing"

	"github.com/josharian/log10"
)

// Alternative implementations,
// partly for testing and partly for fun.

var table10 = [16]uint32{
	0, 1, 10, 100, 1000, 10000,
	100000, 1000000, 10000000, 100000000, 1000000000,
}

func loop(x uint32) int {
	if x == 0 {
		return 1
	}
	ans := 10
	for x < table10[ans&15] {
		ans--
	}
	return ans
}

var tableOff = [32]uint32{
	4:  1<<4 - 10,
	7:  1<<7 - 100,
	10: 1<<10 - 1000,
	14: 1<<14 - 10000,
	17: 1<<17 - 100000,
	20: 1<<20 - 1000000,
	24: 1<<24 - 10000000,
	27: 1<<27 - 100000000,
	30: 1<<30 - 1000000000,
}

func branchless(x uint32) int {
	log2 := bits.Len32(x)
	x += tableOff[log2&31]
	n := bits.Len32(x | 1)
	n *= 77
	n += 256 - 77
	n >>= 8
	return n
}

// Tests.

var names = [...]string{
	"loop",
	"Uint32",
	"branchless",
}

func results(x uint32) [len(names)]int {
	return [...]int{
		loop(x),
		log10.Uint32(x),
		branchless(x),
	}
}

func TestQuick(t *testing.T) {
	var tests []uint32
	// Start with obvious interesting numbers.
	// tests = append(tests, 0, 1, 9, 10, 11 math.MaxUint32-1, math.MaxUint32)
	// Check powers of 10, +/- 1.
	for i := uint64(1); i < math.MaxUint32; i *= 10 {
		tests = append(tests, uint32(i)-1, uint32(i), uint32(i)+1)
	}
	// Check powers of 2, +/- 1.
	for i := uint64(1); i < math.MaxUint32; i *= 2 {
		tests = append(tests, uint32(i)-1, uint32(i), uint32(i)+1)
	}
	tests = append(tests, math.MaxUint32-1, math.MaxUint32, 4294967286)
	for i := range tests {
		a := results(uint32(i))
		for j := 1; j < len(a); j++ {
			if a[0] != a[j] {
				t.Fatalf("log10.%s(%d) = %d, log10.%s(%d) = %d", names[0], i, a[0], names[j], i, a[j])
			}
		}
	}
}

func TestExhaustive(t *testing.T) {
	max := uint64(math.MaxUint32)
	if testing.Short() {
		max = 20000000
	}
	for i := uint64(0); i <= max; i++ {
		if testing.Verbose() && bits.OnesCount64(i) == 1 {
			t.Logf("at %d (%d bits)", i, bits.Len64(i))
		}
		a := results(uint32(i))
		for j := 1; j < len(a); j++ {
			if a[0] != a[j] {
				t.Fatalf("log10.%s(%d) = %d, log10.%s(%d) = %d", names[0], i, a[0], names[j], i, a[j])
			}
		}
	}
}

// Benchmarks.

var sink int

func Benchmark(b *testing.B) {
	b.Run("Predictable", func(b *testing.B) {
		b.Run("loop", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = loop(uint32(i))
			}
		})
		b.Run("Uint32", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = log10.Uint32(uint32(i))
			}
		})
		b.Run("branchless", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = branchless(uint32(i))
			}
		})
	})
}

// Examples.

func ExampleUint32() {
	examples := []uint32{0, 1, 9, 10, 11}
	for _, e := range examples {
		fmt.Printf("log10.Uint32(%d) = %d\n", e, log10.Uint32(e))
	}
	// Output:
	// log10.Uint32(0) = 1
	// log10.Uint32(1) = 1
	// log10.Uint32(9) = 1
	// log10.Uint32(10) = 2
	// log10.Uint32(11) = 2
}

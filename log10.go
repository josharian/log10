// Package log10 calculates log base 10 of an integer.
// This is a simple task.
// This package focuses on performance,
// as there are non-obvious techniques for speed.
package log10

import (
	"math/bits"
)

var table9 = [16]uint32{
	0, 9, 99, 999, 9999, 99999, 999999,
	9999999, 99999999, 999999999, 0xFFFFFFFF,
}

// Uint32 returns the number of digits required to hold the base 10 representation of x.
// As a special case, Uint32(0) == 1.
//
// Note that if you are making a byte buffer for a base 10 representation of x,
// you will often get better results by calling make with a large-enough constant.
// Be sure to benchmark both approaches.
//
// Implementation inspired by
// https://lemire.me/blog/2021/05/28/computing-the-number-of-digits-of-an-integer-quickly/#comment-585476
// and tweaked for the Go compiler.
func Uint32(x uint32) int {
	x |= 1
	log2 := bits.Len32(x)
	n := 9 * log2
	n += 32 - 9
	n >>= 5
	if x > table9[n&15] {
		n++
	}
	return n
}

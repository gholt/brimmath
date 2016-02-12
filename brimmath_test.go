package brimmath

import "testing"

func TestNextHighestPowerOfTwo(t *testing.T) {
	f := NextHighestPowerOfTwo
	for _, vr := range [][]uint64{{1, 1}, {2, 2}, {3, 4}, {4, 4}, {100, 128}, {1000, 1024}, {8191, 8192}, {9223372036854775463, 9223372036854775808}} {
		v, r := vr[0], vr[1]
		if f(v) != r {
			t.Fatal(v, r, f(v))
		}
	}
}

func BenchmarkNextHighestPowerOfTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextHighestPowerOfTwo(uint64(i))
	}
}

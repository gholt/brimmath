// Package brimmath contains math related Go code.
package brimmath

// NextHighestPowerOfTwo returns the next highest power of 2 from v; assuming v
// is > 0. http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
func NextHighestPowerOfTwo(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}

// +build !noasm
// +build sse avx

package vecf64

import "unsafe"

//go:noescape
func sum(a []float64, retVal unsafe.Pointer)

// Sum sums a slice of float64 and returns a float64
func Sum(a []float64) float64 {
	var retVal float64
	sum(a, unsafe.Pointer(&retVal))
	return retVal
}

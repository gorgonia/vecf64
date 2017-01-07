// +build sse avx

package vecf64

/*

IMPORTANT NOTE:

Currently Div does not handle division by zero correctly. It returns a NaN instead of +Inf

*/

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// this file is mainly added to facilitate testing of the ASM code, and that it matches up correctly with the expected results

func TestDiv(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime-1)

	correct := Range(0, niceprime-1)
	for i := range correct {
		correct[i] = correct[i] / correct[i]
	}
	Div(a, a)
	assert.Equal(correct[1:], a[1:])
	assert.Equal(true, math.IsNaN(a[0]), "a[0] is: %v", a[0])

	b := Range(niceprime, 2*niceprime-1)
	for i := range correct {
		correct[i] = a[i] / b[i]
	}

	Div(a, b)
	assert.Equal(correct[1:], a[1:])
	assert.Equal(true, math.IsNaN(a[0]), "a[0] is: %v", a[0])

	/* Weird Corner Cases*/

	for i := 1; i < 65; i++ {
		a = Range(0, i)
		var testAlign bool
		addr := &a[0]
		u := uint(uintptr(unsafe.Pointer(addr)))
		if u&uint(32) != 0 {
			testAlign = true
		}

		if testAlign {
			b = Range(i, 2*i)
			correct = make([]float64, i)
			for j := range correct {
				correct[j] = a[j] / b[j]
			}
			Div(a, b)
			assert.Equal(correct[1:], a[1:])
		}
	}

}

func TestSqrt(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime-1)

	correct := Range(0, niceprime-1)
	for i, v := range correct {
		correct[i] = math.Sqrt(v)
	}
	Sqrt(a)
	assert.Equal(correct, a)

	// negatives
	a = []float64{-1, -2, -3, -4}
	Sqrt(a)

	for _, v := range a {
		if !math.IsNaN(v) {
			t.Error("Expected NaN")
		}
	}

	/* Weird Corner Cases*/
	for i := 1; i < 65; i++ {
		a = Range(0, i)
		var testAlign bool
		addr := &a[0]
		u := uint(uintptr(unsafe.Pointer(addr)))
		if u&uint(32) != 0 {
			testAlign = true
		}

		if testAlign {
			correct = make([]float64, i)
			for j := range correct {
				correct[j] = math.Sqrt(a[j])
			}
			Sqrt(a)
			assert.Equal(correct, a)
		}
	}
}

func TestInvSqrt(t *testing.T) {

	assert := assert.New(t)
	a := Range(0, niceprime-1)

	correct := Range(0, niceprime-1)
	for i, v := range correct {
		correct[i] = 1.0 / math.Sqrt(v)
	}
	InvSqrt(a)
	assert.Equal(correct[1:], a[1:])
	if !math.IsInf(a[0], 0) {
		t.Error("1/0 should be +Inf or -Inf")
	}

	// Weird Corner Cases

	for i := 1; i < 65; i++ {
		a = Range(0, i)
		var testAlign bool
		addr := &a[0]
		u := uint(uintptr(unsafe.Pointer(addr)))
		if u&uint(32) != 0 {
			testAlign = true
		}

		if testAlign {
			correct = make([]float64, i)
			for j := range correct {
				correct[j] = 1.0 / math.Sqrt(a[j])
			}
			InvSqrt(a)
			assert.Equal(correct[1:], a[1:], "i = %d, %v", i, Range(0, i))
			if !math.IsInf(a[0], 0) {
				t.Error("1/0 should be +Inf or -Inf")
			}
		}
	}
}

/* BENCHMARKS */

func _vanillaVecAdd(a, b []float64) {
	for i := range a {
		a[i] += b[i]
	}
}

func BenchmarkVecAdd(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Add(x, y)
	}
}

func BenchmarkVanillaVecAdd(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecAdd(x, y)
	}
}

func _vanillaVecSub(a, b []float64) {
	for i := range a {
		a[i] -= b[i]
	}
}

func BenchmarkVecSub(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Sub(x, y)
	}
}

func BenchmarkVanillaVecSub(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecSub(x, y)
	}
}

func _vanillaVecMul(a, b []float64) {
	for i := range a {
		a[i] *= b[i]
	}
}

func BenchmarkVecMul(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Mul(x, y)
	}
}

func BenchmarkVanillaVecMul(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecMul(x, y)
	}
}

func _vanillaVecDiv(a, b []float64) {
	for i := range a {
		a[i] /= b[i]
	}
}

func BenchmarkVecDiv(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Div(x, y)
	}
}

func BenchmarkVanillaVecDiv(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecDiv(x, y)
	}
}

func _vanillaVecSqrt(a []float64) {
	for i, v := range a {
		a[i] = math.Sqrt(v)
	}
}

func BenchmarkVecSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		Sqrt(x)
	}
}

func BenchmarkVanillaVecSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecSqrt(x)
	}
}

func _vanillaVecInverseSqrt(a []float64) {
	for i, v := range a {
		a[i] = 1.0 / math.Sqrt(v)
	}
}

func BenchmarkVecInvSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		InvSqrt(x)
	}
}

func BenchmarkVanillaVecInvSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecInverseSqrt(x)
	}
}

// +build !avx,!sse

package vecf64

import "math"

func Add(a, b []float64) {
	for i, v := range a {
		a[i] = v + b[i]
	}
}

func Sub(a, b []float64) {
	for i, v := range a {
		a[i] = v - b[i]
	}
}

func Mul(a, b []float64) {
	for i, v := range a {
		a[i] = v * b[i]
	}
}

func Div(a, b []float64) {
	for i, v := range a {
		if b[i] == 0 {
			a[i] = math.Inf(0)
			continue
		}

		a[i] = v / b[i]
	}
}

func Sqrt(a []float64) {
	for i, v := range a {
		a[i] = math.Sqrt(v)
	}
}

func InvSqrt(a []float64) {
	for i, v := range a {
		a[i] = float64(1) / math.Sqrt(v)
	}
}

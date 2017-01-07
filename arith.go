package vecf64

import "math"

// Pow performs a^b elementwise
func Pow(a, b []float64) {
	for i, v := range a {
		switch b[i] {
		case 0:
			a[i] = float64(1)
		case 1:
			a[i] = v
		case 2:
			a[i] = v * v
		case 3:
			a[i] = v * v * v
		default:
			a[i] = math.Pow(v, b[i])
		}
	}
}

// Scale scales all the numbers in a slice by a scalar
func Scale(s float64, a []float64) {
	for i, v := range a {
		a[i] = v * s
	}
}

// DivBy divides all numbers in the slice by a scalar
func DivBy(s float64, a []float64) {
	for i, v := range a {
		a[i] = s / v
	}
}

// Trans translates all thee numbers in the slice by a scalar (slice + scalar)
func Trans(s float64, a []float64) {
	for i, v := range a {
		a[i] = v + s
	}
}

// TransFrom translates all the numbers in a slice from a scalar (scalar - slice)
func TransFrom(s float64, a []float64) {
	for i, v := range a {
		a[i] = s - v
	}
}

// Power performs a^s elementwise
func Power(s float64, a []float64) {
	for i, v := range a {
		a[i] = math.Pow(v, s)
	}
}

// PowerFrom performs s^a eleemntwise
func PowerFrom(s float64, a []float64) {
	for i, v := range a {
		a[i] = math.Pow(s, v)
	}
}

// Max takes two slices, and compares them elementwise. The highest value is put into a
func Max(a, b []float64) {
	if len(a) != len(b) {
		panic("Index error")
	}

	a = a[:len(a)]
	b = b[:len(a)]

	for i, v := range a {
		bv := b[i]
		if bv > v {
			a[i] = bv
		}
	}
}

// Min takes two slices, and compares them elementwise. The lowest value is put into a
func Min(a, b []float64) {
	if len(a) != len(b) {
		panic("Index error")
	}

	a = a[:len(a)]
	b = b[:len(a)]

	for i, v := range a {
		bv := b[i]
		if bv < v {
			a[i] = bv
		}
	}
}

/* REDUCTION RELATED */

func Sum(a []float64) float64 {
	return Reduce(add, float64(0), a...)
}

// MaxOf finds the max of a []float64. it panics if the slice is empty
func MaxOf(a []float64) (retVal float64) {
	if len(a) < 1 {
		panic("Cannot find the max of an empty slice")
	}
	return Reduce(max, a[0], a[1:]...)
}

// MinOf finds the max of a []float64. it panics if the slice is empty
func MinOf(a []float64) (retVal float64) {
	if len(a) < 1 {
		panic("Cannot find the min of an empty slice")
	}
	return Reduce(min, a[0], a[1:]...)
}

// Argmax returns the index of the min in a slice
func Argmax(a []float64) int {
	var f float64
	var max int
	var set bool
	for i, v := range a {
		if !set {
			f = v
			max = i
			set = true

			continue
		}

		// TODO: Maybe error instead of this?
		if math.IsNaN(v) || math.IsInf(v, 1) {
			max = i
			f = v
			break
		}

		if v > f {
			max = i
			f = v
		}
	}
	return max
}

// Argmin returns the index of the min in a slice
func Argmin(a []float64) int {
	var f float64
	var min int
	var set bool
	for i, v := range a {
		if !set {
			f = v
			min = i
			set = true

			continue
		}

		// TODO: Maybe error instead of this?
		if math.IsNaN(v) || math.IsInf(v, -1) {
			min = i
			f = v
			break
		}

		if v < f {
			min = i
			f = v
		}
	}
	return min
}

/* FUNCTION VARIABLES */

var (
	add = func(a, b float64) float64 { return a + b }
	sub = func(a, b float64) float64 { return a - b }
	mul = func(a, b float64) float64 { return a * b }
	div = func(a, b float64) float64 { return a / b }
	mod = func(a, b float64) float64 { return math.Mod(a, b) }

	min = func(a, b float64) float64 {
		if a < b {
			return a
		}
		return b
	}

	max = func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
)

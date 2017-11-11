package vecf64

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 1049 is actually a prime, so it cannot be divisible by any other number
// This is a good way to test that the remainder part of the Add/Sub/Mul/Div/Pow works
const (
	// niceprime = 37
	// niceprime = 1049
	niceprime = 597929
	// niceprime = 1299827 // because sometimes I feel like being an idiot
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime)

	correct := Range(0, niceprime)
	for i := range correct {
		correct[i] = correct[i] + correct[i]
	}

	Add(a, a)
	assert.Equal(correct, a)

	b := Range(niceprime, 2*niceprime)
	for i := range correct {
		correct[i] = a[i] + b[i]
	}

	Add(a, b)
	assert.Equal(correct, a)

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
				correct[j] = b[j] + a[j]
			}
			Add(a, b)
			assert.Equal(correct, a)
		}
	}
}

func TestSub(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime)

	correct := make([]float64, niceprime)
	Sub(a, a)
	assert.Equal(correct, a)

	b := Range(niceprime, 2*niceprime)
	for i := range correct {
		correct[i] = a[i] - b[i]
	}

	Sub(a, b)
	assert.Equal(correct, a)

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
				correct[j] = a[j] - b[j]
			}
			Sub(a, b)
			assert.Equal(correct, a)
		}
	}
}

func TestMul(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime)

	correct := Range(0, niceprime)
	for i := range correct {
		correct[i] = correct[i] * correct[i]
	}
	Mul(a, a)
	assert.Equal(correct, a)

	b := Range(niceprime, 2*niceprime)
	for i := range correct {
		correct[i] = a[i] * b[i]
	}

	Mul(a, b)
	assert.Equal(correct, a)

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
				correct[j] = a[j] * b[j]
			}
			Mul(a, b)
			assert.Equal(correct, a)
		}
	}
}

func TestPow(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}
	b := []float64{0, 1, 2, 3, 4}

	correct := make([]float64, 5)
	for i := range correct {
		correct[i] = math.Pow(a[i], b[i])
	}
	Pow(a, b)
	assert.Equal(t, correct, a)
}

func TestScale(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}

	correct := make([]float64, 5)
	for i := range correct {
		correct[i] = a[i] * 5
	}
	Scale(a, 5)
	assert.Equal(t, correct, a)
}

func TestScaleInv(t *testing.T) {
	a := []float64{0, 1, 2, 4, 6}

	correct := make([]float64, len(a))
	for i := range correct {
		correct[i] = a[i] / 2
	}
	ScaleInv(a, 2)
	assert.Equal(t, correct, a)
}

func TestScaleInvR(t *testing.T) {
	a := []float64{0, 1, 2, 4, 6}
	correct := make([]float64, len(a))
	for i := range correct {
		correct[i] = 2 / a[i]
	}
	ScaleInvR(a, 2)
	assert.Equal(t, correct, a)
}

func TestTrans(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	correct := make([]float64, 4)
	for i := range correct {
		correct[i] = a[i] + float64(1)
	}
	Trans(a, 1)
	assert.Equal(t, correct, a)
}

func TestTransInv(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	correct := make([]float64, 4)
	for i := range correct {
		correct[i] = a[i] - float64(1)
	}
	TransInv(a, 1)
	assert.Equal(t, correct, a)
}

func TestTransInvR(t *testing.T) {
	a := []float64{1, 2, 3, 4}

	correct := make([]float64, len(a))
	for i := range correct {
		correct[i] = float64(1) - a[i]
	}
	TransInvR(a, 1)
	assert.Equal(t, correct, a)
}

func TestPowOf(t *testing.T) {
	a := []float64{1, 2, 3, 4}

	correct := make([]float64, len(a))
	for i := range correct {
		correct[i] = math.Pow(a[i], 5)
	}
	PowOf(a, 5)
	assert.Equal(t, correct, a)
}

func TestPowOfR(t *testing.T) {
	a := []float64{1, 2, 3, 4}

	correct := make([]float64, len(a))
	for i := range correct {
		correct[i] = math.Pow(5, a[i])
	}
	PowOfR(a, 5)
	assert.Equal(t, correct, a)
}

func TestMax(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}
	b := []float64{5, 4, 2, 2, 1}
	correct := []float64{5, 4, 2, 3, 4}

	Max(a, b)
	assert.Equal(t, correct, a)

	b = []float64{2}
	f := func() {
		Max(a, b)
	}
	assert.Panics(t, f)
}

func TestMin(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}
	b := []float64{5, 4, 2, 2, 1}
	correct := []float64{0, 1, 2, 2, 1}

	Min(a, b)
	assert.Equal(t, correct, a)

	b = []float64{2}
	f := func() {
		Min(a, b)
	}
	assert.Panics(t, f)
}

func TestSum(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}
	correct := float64(10)
	got := Sum(a)
	if correct != got {
		t.Errorf("Expected %f. Got %v instead", correct, got)
	}
}

func TestMaxOf(t *testing.T) {
	a := []float64{0, 1, 2, 1, 0}
	correct := float64(2)
	got := MaxOf(a)
	if got != correct {
		t.Errorf("Expected %f. Got %v instead", correct, got)
	}

	a = []float64{}
	f := func() {
		MaxOf(a)
	}
	assert.Panics(t, f, "Expected panic when empty slice passed into MaxOf")
}

func TestMinOf(t *testing.T) {
	a := []float64{0, 1, 2, 1, 0}
	correct := float64(0)
	got := MinOf(a)
	if got != correct {
		t.Errorf("Expected %f. Got %v instead", correct, got)
	}

	a = []float64{}
	f := func() {
		MinOf(a)
	}
	assert.Panics(t, f, "Expected panic when empty slice passed into MinOf")
}

func TestArgmax(t *testing.T) {
	a := []float64{0, 1, 2, 34, 5}
	correct := 3
	got := Argmax(a)
	if got != correct {
		t.Errorf("Expected argmax to be %v. Got %v instead", correct, got)
	}

	a = []float64{math.Inf(-1), 2, 3, 4}
	correct = 3
	got = Argmax(a)
	if got != correct {
		t.Errorf("Expected argmax to be %v. Got %v instead", correct, got)
	}

	a = []float64{math.Inf(1), 2, 3, 4}
	correct = 0
	got = Argmax(a)
	if got != correct {
		t.Errorf("Expected argmax to be %v. Got %v instead", correct, got)
	}

	a = []float64{1, math.NaN(), 3, 4}
	correct = 1
	got = Argmax(a)
	if got != correct {
		t.Errorf("Expected argmax to be %v. Got %v instead", correct, got)
	}
}

func TestArgmin(t *testing.T) {
	a := []float64{0, 1, 2, -34, 5}
	correct := 3
	got := Argmin(a)
	if got != correct {
		t.Errorf("Expected argmin to be %v. Got %v instead", correct, got)
	}

	a = []float64{math.Inf(-1), 2, 3, 4}
	correct = 0
	got = Argmin(a)
	if got != correct {
		t.Errorf("Expected argmin to be %v. Got %v instead", correct, got)
	}

	a = []float64{math.Inf(1), 2, 3, 4}
	correct = 1
	got = Argmin(a)
	if got != correct {
		t.Errorf("Expected argmin to be %v. Got %v instead", correct, got)
	}

	a = []float64{1, math.NaN(), 3, 4}
	correct = 1
	got = Argmin(a)
	if got != correct {
		t.Errorf("Expected argmin to be %v. Got %v instead", correct, got)
	}
}

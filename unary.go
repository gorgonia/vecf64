package vecf64

import "math"

func TanhRecv(a, retVal []float64) {
	retVal = retVal[:len(a)]
	for i, v := range a {
		retVal[i] = math.Tanh(v)
	}
}

func Tanh(a []float64) {
	for i, v := range a {
		a[i] = math.Tanh(v)
	}
}

func ExpRecv(a, retVal []float64) {
	retVal = retVal[:len(a)]
	for i, v := range a {
		retVal[i] = math.Exp(v)
	}
}

func Exp(a []float64) {
	for i, v := range a {
		a[i] = math.Exp(v)
	}
}

func LogRecv(a, retVal []float64) {
	retVal = retVal[:len(a)]
	for i, v := range a {
		retVal[i] = math.Log(v)
	}
}

func Log(a []float64) {
	for i, v := range a {
		a[i] = math.Log(v)
	}
}

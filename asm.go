// +build sse avx

package vecf64

func Add(a, b []float64)
func Sub(a, b []float64)
func Mul(a, b []float64)
func Div(a, b []float64)

func Sqrt(a []float64)
func InvSqrt(a []float64)

/*
func Pow(a, b []float64)
*/

/*
func Scale(s float64, a []float64)
func ScaleFrom(s float64, a []float64)
func Trans(s float64, a []float64)
func TransFrom(s float64, a []float64)
func Power(s float64, a []float64)
func PowerFrom(s float64, a []float64)
*/

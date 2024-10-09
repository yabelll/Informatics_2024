package lab

import (
	"math"
)

func Calc(b float64, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func TaskA(b float64, Xn float64, Xk float64, delX float64) []float64 {
	values := []float64{}
	for x := Xn; x <= Xk; x += delX {
		values = append(values, Calc(b, x))
	}
	return values
}

func TaskB(b float64, x [5]float64) []float64 {
	values := []float64{}
	for _, value := range x {
		values = append(values, Calc(b, value))
	}
	return values
}
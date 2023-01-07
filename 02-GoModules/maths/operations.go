package maths

import "math"

func Sum(a, b int) int {
	return a + b
}

func Subs(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) float64 {
	return float64(a) / float64(b)
}

func Pot(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

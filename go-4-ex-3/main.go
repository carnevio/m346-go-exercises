package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) []float64 {
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		return []float64{}
	}
	if D == 0 {
		x := -b / (2 * a)
		return []float64{x}
	}
	sqrtD := math.Sqrt(D)
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)
	return []float64{x1, x2}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println("computeQuadraticFormula(3,4,1) =", computeQuadraticFormula(3, 4, 1)) // two solutions: [-0.3333333333333333 -1]
	// D = 0
	fmt.Println("computeQuadraticFormula(2,4,2) =", computeQuadraticFormula(2, 4, 2)) // one solution: [-1]
	// D < 0
	fmt.Println("computeQuadraticFormula(3,4,2) =", computeQuadraticFormula(3, 4, 2)) // no solution: []
}

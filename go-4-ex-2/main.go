package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a float64
	b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return computeHypotenuse(s.a, s.b)
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println("computeHypotenuse(3, 4) =", computeHypotenuse(3, 4))   // 5
	fmt.Println("computeHypotenuse(5, 12) =", computeHypotenuse(5, 12)) // 13
	fmt.Println("computeHypotenuse(1, 1) =", computeHypotenuse(1, 1))   // 1.4142135623730951
	fmt.Println("computeHypotenuse(0, 5) =", computeHypotenuse(0, 5))   // 5

	s1 := ShortSides{a: 3, b: 4}
	fmt.Println("ShortSides{3,4}.Hypotenuse() =", s1.Hypotenuse()) // 5

	s2 := ShortSides{a: 5, b: 12}
	fmt.Println("ShortSides{5,12}.Hypotenuse() =", s2.Hypotenuse()) // 13

	s3 := ShortSides{a: 1, b: 1}
	fmt.Println("ShortSides{1,1}.Hypotenuse() =", s3.Hypotenuse()) // 1.4142135623730951
}

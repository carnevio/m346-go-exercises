package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) {
	y := (gotPoints/maxPoints)*5 + 1
	fmt.Println(y)

}

func main() {
	// TODO: call the function computeGrade
	computeGrade(17.5, 28.0) // 4.125
	computeGrade(20.5, 30.0) // 4.416667
	computeGrade(28.0, 28.0) // 6
}

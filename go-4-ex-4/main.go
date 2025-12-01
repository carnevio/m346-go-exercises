package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32.0
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Println("convertCelsiusToFahrenheit(0) =", convertCelsiusToFahrenheit(0))     // 32
	fmt.Println("convertCelsiusToFahrenheit(100) =", convertCelsiusToFahrenheit(100)) // 212
	fmt.Println("convertCelsiusToFahrenheit(-40) =", convertCelsiusToFahrenheit(-40)) // -40
	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println("convertFahrenheitToCelsius(32) =", convertFahrenheitToCelsius(32))   // 0
	fmt.Println("convertFahrenheitToCelsius(212) =", convertFahrenheitToCelsius(212)) // 100
	fmt.Println("convertFahrenheitToCelsius(-40) =", convertFahrenheitToCelsius(-40)) // -40
}

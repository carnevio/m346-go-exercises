package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName = "Nevio"
	var lastName  = "Carcangiu"

	var dayOfBirth = 1
	var monthOfBirth = 8
	var yearOfBirth = 2008

	var numberOfSiblings = 1

	var heightInMeters = 1.80

	var zodiacSign = '\u264C'


	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}

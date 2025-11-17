package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Nevio",
			LastName:  "Carcangiu",
		},
		BirthDate: BirthDate{
			Day:   1,
			Month: 8,
			Year:  2008,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       '\u264C',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}

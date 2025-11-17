package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	classA := Class{
		Name: "1A",
		Students: []Student{
			{"Anna", "Müller"},
			{"Luca", "Schneider"},
			{"Maya", "Fischer"},
		},
	}

	classB := Class{
		Name: "2B",
		Students: []Student{
			{"Noah", "Meier"},
			{"Sofia", "Keller"},
			{"Leon", "Zimmer"},
		},
	}

	classC := Class{
		Name: "3C",
		Students: []Student{
			{"Emma", "Becker"},
			{"Luis", "Hoffmann"},
			{"Mia", "Schulz"},
		},
	}
	modules := map[int][]Class{
		104: {classA, classB}, // z.B. Programmieren I
		117: {classB, classC}, // z.B. Datenbanken
		346: {classA, classC}, // z.B. Cloud-Lösungen konzipieren und realisieren
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}

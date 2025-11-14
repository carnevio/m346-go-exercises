package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Zufallszahlengenerator initialisieren

	eyes := rand.Intn(6) + 1
	when := time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?

	//$go run go-1-ex-3/main.go >out.txt 2>&1
	//$cat out.txt
	//$the dice shows 1 eyes
	//$the dice was rolled at 2025-11-03 15:12:07.4377596 +0100 CET m=+0.000512401

	// Aufrufbeispiel (Windows, aus Projekt-Root):
	// go run .\go-1-ex-3\main.go > eyes.txt 2> dice.log
	// stdout -> eyes.txt  (die Augenzahl)
	// stderr -> dice.log  (der Zeitstempel)
	// Anzeigen: type eyes.txt && type dice.log

	// go run ex3/main.go TODO
}

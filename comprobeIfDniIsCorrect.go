package main

import (
	"fmt"
	"strconv"
)

var dnis = []string{"48204767X", "48204766F", "02321331X", "02323113X"}
func main() {
	comprobeDnis()
}

func comprobeDnis() {
	for _,dni := range dnis {
		var letter = dni[8:9]

		correctLetter := calculateLetterOfDni(dni)
		if letter == correctLetter {
			fmt.Println("DNI correcto")
			fmt.Println(dni, correctLetter)
		} else {
			fmt.Println("DNI incorrecto")
			fmt.Println(dni, correctLetter)
		}
	}
}

func calculateLetterOfDni(dni string) string {
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var dniNumber = dni[:8]
	//conevertir el dni a numero
	var dniNumberInt, _ = strconv.Atoi(dniNumber)
	dniLetter := letters[dniNumberInt%23]
	return dniLetter
}
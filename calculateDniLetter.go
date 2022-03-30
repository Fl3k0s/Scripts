package main

import "fmt"

var dni = 73434769

func main() {
	letter := calculateDniLetter()
	fmt.Printf("The letter of dni %d is %s\n", dni, letter)
}

//calculateDniLetter returns the letter of the dni
func calculateDniLetter() string {
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	return letters[dni % 23]
}

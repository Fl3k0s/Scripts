package main

import (
	"fmt"
)

func main() {
	var str = "Hello World"
	var count = 0
	count = len(str)
	fmt.Printf("the length of '%s' is: %d", str, count)
}

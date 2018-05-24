package main

import (
	"fmt"
)

func main() {
	// Slices pointers are diferents from a struct
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

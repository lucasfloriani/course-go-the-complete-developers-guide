package main

import (
	"fmt"
)

func main() {
	// Create a slice of integers
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Loop throught the numbers
	for _, number := range numbers {
		// Check if the number is even and print
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else { // else print
			fmt.Println(number, "is odd")
		}
	}
}

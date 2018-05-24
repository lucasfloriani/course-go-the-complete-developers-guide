package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// First way to create a variable of type person
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)

	// Second way to create a variable of type person
	lucas := person{firstName: "Lucas", lastName: "Floriani"}
	fmt.Println(lucas)
	fmt.Println(lucas.firstName)
	fmt.Println(lucas.lastName)
}

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	/*
	 * == Zero value ==
	 * When a variable is created without value assign
	 *
	 * string == ""
	 * int == 0
	 * float == 0
	 * bool == false
	 **/

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

	// Third way to create a variable of type person
	var roberto person
	fmt.Println(roberto)
	fmt.Println(roberto.firstName)
	fmt.Println(roberto.lastName)
	// %+v will print all the diferent fields
	// names and their values from roberto
	fmt.Printf("%+v\n", roberto)
	roberto.firstName = "Roberto"
	roberto.lastName = "Rural"
	fmt.Printf("%+v", roberto)
}

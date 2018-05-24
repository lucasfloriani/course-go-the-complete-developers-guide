package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// &jim give the memory address of the value
	// this variable is pointing at
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// *person declares the type of variable the receiver is,
// a person pointer
func (pointerToPerson *person) updateName(newFirstName string) {
	// *person give me the value this
	// memory address is pointing ate
	(*pointerToPerson).firstName = newFirstName
}

// Turn "address" into "value" by "*address"
// Turn "value" into "address" by "&value"

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // Same as "contactInfo contactInfo"
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

	// Copy all values from jim and pass to p
	jim.print()
	// Copy all values from jim and pass to p
	// Dont change the name of type person
	jim.updateName("Jimmy")
	// Copy all values from jim and pass to p
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

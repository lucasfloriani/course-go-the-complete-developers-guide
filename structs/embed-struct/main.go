package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Custom type in a custom type
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v\n", jim)
}

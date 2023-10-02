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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 12345,
		},
	}

	alex.updateName("Alexander")
	alex.print()
}

// print prints out the values of the person
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// updateName updates the first name of person
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

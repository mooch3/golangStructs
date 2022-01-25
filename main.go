package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}
// embedded struct
type person struct {
	firstName string
	lastName string
	contactInfo
}



func main() {
	// jim := person{ 
	// firstName: "Jim",
	// lastName: "Bund",	
	// contactInfo: contactInfo {
	// 	email: "123@jim.com",
	// 	zip: 94327,
	// 	},
	// }
	// // get memory address
	// // can use a short cut to change variable of type person into pointer of type person
	// jim.updateName("Carl")
	// jim.print()

	name := "Billy"

	fmt.Println(*&name)
}
// *person means we are working with type pointer to person (type description)
// *pointerToPerson is an operator and means we want to manipulate the value
func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func (p person) print() {
fmt.Printf("%+v", p)
}
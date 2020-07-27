package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //we can remove contact name here and just keep contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex:=person{"Alex","Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{ //contactInfo:contactInfo{}
			email:   "@gmail.com",
			zipCode: 273014,
		},
	}

	//	fmt.Printf("%+v", jim)

	//jimPoint := &jim
	//jimPoint.updateName("Jimmy")

	jim.updateName("Jimmy") //36,37 will work same as this
	jim.print()

	mySlice := []string{"Hi", "How", "Are", "You"}
	update(mySlice)
	fmt.Println(mySlice)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (ppointerToPerson *person) updateName(newFirstName string) {
	(*ppointerToPerson).firstName = newFirstName
}

func update(s []string) {
	s[0] = "bye"
}

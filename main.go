package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("myfile123")
	// cards := newDeckFromFile("mayfile124")
	// cards.print()

	//cards := newDeck()
	//cards.shuffle()
	//cards.print()

	// alex := person{firstName: "Mehran", lastName: "Janfeshan"}
	//var alex person
	//alex.lastName = "Janfeshan"
	//alex.firstName = "Mehran"
	//alex.contact.email = "janfeshan.mehran@gmail.com"
	//alex.contact.zipCode = 342345234

	alext := person{name: "Mehran", lastname: "Janfeshan",
		contactInfo: contactInfo{
			email:   "jan@jin.com",
			zipCode: "23423",
		}}

	alext.updateName("Jafar")
	alext.print()
}

func (p *person) updateName(name string) {
	p.name = name
}

func (p person) print() {
	fmt.Println(p)
}

type person struct {
	name     string
	lastname string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

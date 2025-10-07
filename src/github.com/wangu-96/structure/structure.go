package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) getDetails() string {
	return "User details: \nFirst Name: " + p.firstName + "\nLast Name: " + p.lastName + "\nAge: " + strconv.Itoa(p.age)
}

func (p *Person) setAge(newAge int) {
	p.age = newAge
}

func main() {

	person1 := Person{"Wangu", "Ngalati", 28}

	person1.setAge(30)

	personDetailsSeperate := "User details: \nFirst Name: " + person1.firstName + "\nLast Name: " + person1.lastName + "\nAge: " + strconv.Itoa(person1.age)

	fmt.Println(personDetailsSeperate)

	fmt.Println("User details", person1)

	fmt.Println()

	fmt.Println("Reciever or getter", person1.getDetails())

}

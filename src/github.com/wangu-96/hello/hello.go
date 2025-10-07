package main

import "fmt"

var name = "Wangu Ngalati"

func main() {
	fmt.Println("Hello, World!")

	email, password := "wangungalati@gmail.com", "password123"
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)
	fmt.Println("Name:", name)

	fmt.Printf("Type of email: %T\n", email)
	fmt.Printf("Type of password: %T\n", password)
	fmt.Printf("Type of name: %T\n", name)

}

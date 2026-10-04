package main

import "fmt"

func main() {
	fmt.Println("Welcome to the application")

	var name string
	fmt.Println("Please enter your name...")
	fmt.Scanln(&name)

	fmt.Println("Hello, ", name)

}

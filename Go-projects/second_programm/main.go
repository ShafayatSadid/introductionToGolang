package main

import "fmt"

func printWelcome() {
	fmt.Println("Welcome to the application")
}

func getUsersName() string {
	var name string
	fmt.Println("Please enter your name...")
	fmt.Scanln(&name)

	return name
}

func getNumbers() (int, int) {
	var num1 int
	fmt.Println("please enter first number...")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Println("Enter second number...")
	fmt.Scanln(&num2)

	return num1, num2
}

func Add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func display(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("Sum of your given numbers = ", sum)
}

func finalMessage() {
	fmt.Println("Thanks for using our application")
	fmt.Println("Goodby")
}

func main() {
	printWelcome()

	usersName := getUsersName()
	
	num1, num2 := getNumbers()
	sum := Add(num1, num2)

	display(usersName, sum)
	finalMessage()
}

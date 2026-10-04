package main

import "fmt"

func sum(a int, b int) int {
	sum := a + b
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	mul := num1 * num2

	return sum, mul
}

func printSomething (){
	fmt.Println("Education must be free")
}

func sayHello(name string){
	fmt.Println("Welcome to the Golang course, ", name)
}

func main() {

	sayHello("Shafayat")
}

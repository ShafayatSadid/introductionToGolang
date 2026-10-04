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

func main() {

	a := 10
	b := 5

	sum, mul := getNumbers(a, b)

	fmt.Println("the sum is = ", sum)
	fmt.Println("the multiplication is = ", mul)
}

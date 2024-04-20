package main

import "fmt"

func factorial(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result = result * i
	}
	return result
}

func factorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialRecursive(number-1)
	}
}
func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorialRecursive(3))
}

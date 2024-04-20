package main

import "fmt"

func prima(number int) {

	if number == 3 {
		fmt.Println("prima")
	} else if number%3 != 0 && number%2 != 0 {
		fmt.Println("prima")
	} else {
		fmt.Println("tidak prima")
	}
}

func main() {

	prima(7)
}

package main

import "fmt"

func fizzbuzz(number int) {
	for i := 1; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}
	}
}

func main() {
	fizzbuzz(20)
}

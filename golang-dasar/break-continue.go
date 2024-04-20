package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("index ke-", i)
	}
}

package main

import "fmt"

func linearsearch(arr []int, number int) {
	len := len(arr)

	for i := 0; i < len; i++ {
		if arr[i] == number {
			fmt.Println("found at index -", i)
			return
		}
	}
	fmt.Println("Not found")
}

func main() {
	array := []int{12, 13, 90, 1, 5, 2, 3, 10}
	linearsearch(array, 1)
}

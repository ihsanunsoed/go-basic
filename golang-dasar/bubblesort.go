package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr)
		}
	}
}

func main() {
	// Contoh penggunaan
	array := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Array sebelum diurutkan:", array)

	bubbleSort(array)

	fmt.Println("Array setelah diurutkan:", array)
}

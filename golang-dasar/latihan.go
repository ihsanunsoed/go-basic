package main

import "fmt"

func recursive(number int) int {
	if number == 0 || number == 1 {
		return 1
	} else {
		return number * recursive(number-1)
	}
}

func ganjil(number int) {
	if number%2 != 0 {
		fmt.Println("ganjil", number)
	}
}

func bubblesort(arr []int) {
	lenght := len(arr)

	for i := 0; i < lenght-1; i++ {
		for j := 0; j < lenght-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func binarysearch(arr []int, target int) int {
	min, max := 0, len(arr)-1

	for min <= max {
		mid := (min + max) / 2
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

func main() {
	recursive(10)
	ganjil(2)
	fmt.Println("hasil", recursive(5))

	arr := []int{1, 2, 5, 8, 10, 7, 9, 11}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bubblesort(arr)
	find := binarysearch(arr2, 6)
	if find != -1 {
		fmt.Println("found at index", find)
	}

}

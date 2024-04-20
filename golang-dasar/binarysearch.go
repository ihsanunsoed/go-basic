package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // elemen ditemukan, kembalikan indeks
		} else if arr[mid] < target {
			low = mid + 1 // cari di sebelah kanan mid
		} else {
			high = mid - 1 // cari di sebelah kiri mid
		}
	}

	return -1 // elemen tidak ditemukan
}

func main() {
	// Contoh penggunaan
	sortedArray := []int{11, 12, 22, 25, 34, 64, 90}
	target := 12

	result := binarySearch(sortedArray, target)

	if result != -1 {
		fmt.Printf("Elemen %d ditemukan pada indeks %d.\n", target, result)
	} else {
		fmt.Printf("Elemen %d tidak ditemukan dalam array.\n", target)
	}
}

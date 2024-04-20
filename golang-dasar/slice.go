package main

import "fmt"

func main() {

	names := [...]string{"rizal", "novi", "ihsan", "buah", "sayur"}
	//cara deklarasi slice
	slices1 := names[1:3]

	fmt.Println(slices1)

	slices2 := names[:5]

	fmt.Println(slices2)

	var slice3 []string = names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "jumat")
	fmt.Println(daySlice2)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "ihsan"
	newSlice[1] = "punta"
	// newSlice[2] = "Eko" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko", "eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	iniSlice2 := append(iniSlice, 1, 2, 3, 4, 5)

	fmt.Println(iniSlice2)

}

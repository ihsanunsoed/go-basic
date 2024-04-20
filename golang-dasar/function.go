package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}
func tambah(angka1 int, angka2 int) {
	angka3 := angka1 + angka2
	fmt.Println(angka3)
}

func kali(angka1 int, angka2 int) int {
	angka3 := angka1 * angka2
	return angka3
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return firstName, middleName, lastName
}

func main() {
	tambah(1, 2)
	fmt.Println(kali(5, 5))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// firstName, _ := getFullName()
	// fmt.Println(firstName)

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

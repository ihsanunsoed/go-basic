package main

import "fmt"

//any adalah interface kosong interface{} fungsi yang mengembalikan nilai interface kosong atau any bisa mengembalikan semua tipe data
func Ups() any {
	//return 1
	//return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}

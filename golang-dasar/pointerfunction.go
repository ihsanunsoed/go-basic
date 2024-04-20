package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountrtyToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	var address Address = Address{}
	ChangeCountrtyToIndonesia(&address)
	fmt.Println(address)
}

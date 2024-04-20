package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// address1 := Address{
	// 	City:     "Bandung",
	// 	Province: "Jawa Barat",
	// 	Country:  "Indonesia",
	// }

	// fmt.Println(address1)

	// address2 := &address1
	// address2.City = "Purbalingga"

	var address1 Address = Address{City: "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia"}

	var address2 *Address = &address1
	address2.City = "purbalingga"
	fmt.Println(address1)
	fmt.Println(address2)
}

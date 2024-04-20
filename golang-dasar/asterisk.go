package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandung

	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} //tidak merubah address 1
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} //merubah semua yang direfference address2
	fmt.Println(address1)
	fmt.Println(address2)
}

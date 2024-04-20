package main

import "fmt"

//struct merupakan blueprint data, atau seperti class pada java dan php
type Customer struct {
	Name    string
	Address string
	Number  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("halo", name, "my name is", customer.Name)
}

func main() {
	//membuat objek dari struct
	var ihsan Customer
	ihsan.Name = "ihsan"
	ihsan.Address = "telaga harapan"
	// ihsan.Number = 6281382221433

	fmt.Println(ihsan)
	//cara lain membuat objek dari struct tetapi urutan harus sama
	rijal := Customer{
		Name:    "Rizal",
		Address: "Mewek",
		Number:  12313,
	}

	fmt.Println(rijal)

	rully := Customer{Name: "rully"}
	rully.sayHello("natash")
}

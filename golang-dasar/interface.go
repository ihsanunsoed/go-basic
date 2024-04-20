package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct {
	Name string
}

//jika suatu struct memiliki field yang sama dengan suatu interface makan otomatis interface akan diimplementasikan pada struct tersebut
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	ihsan := Person{
		Name: "natashhh",
	}
	SayHello(ihsan)
	animal := Animal{Name: "Kucing"}
	SayHello(animal)

}

package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "ihsan",
		"alamat": "telaga harapan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["alamat"])

	book := make(map[string]string)
	book["title"] = "titanic"
	book["author"] = "jk rowling"
	book["tes"] = "salah"

	fmt.Println(book)
	delete(book, "tes")

	fmt.Println(book)
}

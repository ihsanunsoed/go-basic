package main

import "fmt"

func main() {

	var names [5]string
	names[0] = "rijal"
	names[1] = "ihsan"
	names[2] = "aku"
	names[3] = "amar"
	names[4] = "daffa"

	fmt.Println(names)
	fmt.Println(len(names))

	var class = [...]int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(class)
	fmt.Println(len(class))

}

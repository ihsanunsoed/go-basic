package main

import "fmt"

func random() any {
	return "true"
}

func main() {
	//fungsi ini mengembalikan interface kosong yang mana bisa mengembalikan tipe data apapun
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	//type assertion dengan switch case
	//type assertion adalah merubah tipe data dari kembalian fungsi
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

package main

import "fmt"

func main() {

	for counter := 5; counter > 0; counter-- {
		fmt.Println("")
		for counter2 := 0; counter2 < counter; counter2++ {
			fmt.Print("*")
		}
	}

	names := [5]string{
		"rizal",
		"ihsan",
		"aku",
		"kamu",
		"kita",
	}

	//2 nama variable didepan bebas tapi yang pertama akan menjadi index dan yang kedua menjadi valuenya
	for index, name := range names {
		fmt.Println("index ke-", index, "nama :", name)
	}

	//jika kita hanya ingin menggunakan value, kita bisa ubah variable indexny jadi _
	for _, name := range names {
		fmt.Println("nama :", name)
	}

}

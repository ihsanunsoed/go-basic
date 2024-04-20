package main

import (
	"fmt"
	"golang-dasar/database"
)

func main() {
	var koneksi = database.GetDatabase()

	fmt.Println(koneksi)
}

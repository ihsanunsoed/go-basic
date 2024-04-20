package main

import "fmt"

func logging() {
	fmt.Println("fungsi selesai dijalankan")
}
func runApplication() {
	//defer function adalah function yang dipanggil ketika function selesai dijalankan walaupun functionnya error deffer function akan tetap di eksekusi
	defer logging()
	fmt.Println("running")

}
func main() {
	runApplication()
}

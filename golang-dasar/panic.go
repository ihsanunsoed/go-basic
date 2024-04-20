package main

import "fmt"

func endApp() {
	fmt.Println("app closed")
	//ketika menggunakan recover walaupun ada panic aplikasi akan tetap berjalan
	message := recover()
	fmt.Println("error :", message)
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("running app...")
	if error {
		panic("ups error")
	}
	fmt.Println("closing app...")
}
func main() {
	runApp(false)
}

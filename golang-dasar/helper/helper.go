package helper

import "fmt"

var version = "1.0.0"            //karena huruf depannya kecil variabel ini tidak bisa di akses dri luar package
var Application = "golang/dasar" // ini bisa krena huruf depannya besar

// ini tidak bisa
func sayGoodbye(name string) string {
	return ("bye " + name)
}

func SayHello(name string) string {
	return ("halo " + name)
}

//walaupun tidak bisa diakses dri packake lain, variabel dan fungsinya masih bisa diakses dri package yang sama

func contoh() {
	fmt.Println(version)
	sayGoodbye("san")
}

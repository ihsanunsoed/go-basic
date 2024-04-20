package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	result := helper.SayHello("San")
	fmt.Println(helper.SayHello("Natashhhh") + " " + result)

	fmt.Println(helper.Application)
}

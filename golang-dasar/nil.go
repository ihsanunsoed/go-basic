package main

import "fmt"

func NewMap(name string, address string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name":    name,
			"address": address,
		}
	}
}

func main() {
	data := NewMap("Eko", "telaga")

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"], data["address"])
	}
}

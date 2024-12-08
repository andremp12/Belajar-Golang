package main

import "fmt"

func main() {
	name := "Aji"

	if name == "Andre" {
		fmt.Println("Hello, ", name)
	} else if name == "Aji" {
		fmt.Println("Hello, ", name)
	} else {
		fmt.Println("Salah")
	}

	if length := len(name); length > 4 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}

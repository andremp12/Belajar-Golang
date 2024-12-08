package main

import "fmt"

func main() {
	name := "Andre"

	//Switch Menggunakan Kondisi
	switch name {
	case "Andre":
		fmt.Println("Hello Andre")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hello, Boleh Kenalan ?")
	}

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("\nNama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch Tanpa Kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("\nNama terlalu panjang")
	case length > 4:
		fmt.Println("\nNama lumayan panjang")
	default:
		fmt.Println("\nNama sudah benar")
	}
}

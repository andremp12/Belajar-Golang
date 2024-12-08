package main

import "fmt"

func main() {
	var nama [3]string

	nama[0] = "Andre"
	nama[1] = "Zakki"
	nama[2] = "Aji"

	fmt.Println(nama[0])

	var number = [4]int{
		30,
		12,
		15,
	}

	fmt.Println(number[2])
	fmt.Println("Panjang Array Nama :", len(nama))
	fmt.Println("Panjang Array Number :", len(number))
}

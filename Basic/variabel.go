package main

import "fmt"

func main() {
	//Deklarasi Variabel

	//String
	var nama1 string
	var nama2 = "Jose"
	nama3 := "Andi"
	var (
		nama4 = "Ridho"
		nama5 = "Angga"
	)

	nama1 = "Andre Mahesa"

	fmt.Println(nama1)
	fmt.Println(nama2)
	fmt.Println(nama3)
	fmt.Println(nama4)
	fmt.Println(nama5)

	//Integer
	var age int8 = 21
	var (
		number1 int16 = 100
		number2 int32 = 10000
	)

	fmt.Println(age)
	fmt.Println(number1)
	fmt.Println(number2)
}

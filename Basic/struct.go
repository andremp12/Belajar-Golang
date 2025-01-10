package main

import "fmt"

/*
Struct merupakan sebuah template data yg digunakan untuk menggabungkan
beberapa jenis tipe data dalam kesatuan

Data di struct disimpan dalam field

Jadi, Struct merupakan kumpulan dari field
*/

/*
Struct tidak bisa langsung digunakan,
Namun kita bisa membuat data/object dari struct yg telah dibuat
*/

type Customer struct {
	name    string
	address string
	age     int
	jenkel  string
}

/*func sayHi(customer Customer) {
	fmt.Println("Hello", customer.name)
}*/

// Struct parameter
func (customer Customer) sayHi() {
	fmt.Println("Hello", customer.name)
}

func main() {
	var aji Customer
	aji.name = "Ajiii"
	aji.address = "Padang"
	aji.age = 22
	aji.jenkel = "Laki-Laki"

	/*fmt.Println(aji)
	fmt.Println(aji.address)

	ridho := Customer{
		name:    "Ridho",
		address: "Cirebon",
		age:     30,
		jenkel:  "Laki-Laki",
	}
	fmt.Println(ridho)

	budi := Customer{"Budi", "Jakarta", 20, "Laki-Laki"}
	fmt.Println(budi)*/

	//sayHi(aji)
	aji.sayHi()
}

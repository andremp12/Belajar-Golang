package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Padang", "Sumbar", "Indonesia"}

	//Pass by Value
	//address2 := address1
	//
	//address2.City = "Solok"
	//
	//fmt.Println(address1)
	//fmt.Println(address2)

	//Pointer --> Pass by Reference
	address2 := &address1
	address3 := &address1
	address2.City = "Solok Selatan"

	fmt.Println(address1) //adress 1 ikut berubah
	fmt.Println(address2)

	//address2 = &Address{"Pekanbaru", "Riau", "Indonesia"}
	//
	//fmt.Println(address1) //Tidak akan merubah address1, karna address2 di assign ulang dan akan membuat data baru
	//fmt.Println(address2)
	//fmt.Println(address3)

	*address2 = Address{"Pekanbaru", "Riau", "Indonesia"}

	fmt.Println(address1) // semua variabel yg pointer ke data yang sama, yaitu ke data address1, nilainya berubah
	fmt.Println(address2)
	fmt.Println(address3)

	//function new, membuat pointer dengan data kosong
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

}

package main

import "fmt"

type Employee struct {
	Name, Province, Country string
	Age                     int
}

//func ChangeCountryToIndonesia(employee Employee) {
//	employee.Country = "Indonesia"
//}

// pointer di function : Lebih disarankan ketika struct memiliki field,
// atau data yang banyak, sehingga dapat menghemat memory
func ChangeCountryToIndonesia(employee *Employee) {
	employee.Country = "Indonesia"
}

func main() {
	employee := Employee{
		Name:     "Andre",
		Province: "Sumbar",
		Country:  "",
		Age:      20,
	}

	ChangeCountryToIndonesia(&employee)

	fmt.Println(employee) // data country employee akan berubah
}

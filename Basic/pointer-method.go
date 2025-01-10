package main

import "fmt"

type Man struct {
	Name string
	Age  int
}

//func (man Man) Married() {
//	man.Name = "Mr." + man.Name
//}

// Rata-rata orang menggunakan pointer saat membuat struct function
func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	Andre := Man{
		Name: "Andre",
		Age:  20,
	}

	// Data Andre tidak berubah, karena pada struct function belum memiliki pointer
	//Andre.Married()
	//fmt.Println(Andre)

	// Data Amdre berubah, setelah membuat pointer pada struct function
	Andre.Married()
	fmt.Print(Andre)
}

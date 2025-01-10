package main

import "fmt"

func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := months[4:7]
	fmt.Println(slice1)

	//function slice
	fmt.Println("Lenght Slice :", len(slice1))
	fmt.Println("Capacity Slice :", cap(slice1))

	/*months[5] = "Kustom"
	fmt.Println("Nilai slice diubah :", slice1)*/

	/*slice1[0] = "Kustom2"
	fmt.Println("Nilai Array month diubah :", months)*/

	slice2 := months[9:]
	slice3 := append(slice2, "Baru") /*function append berfungsi ketika panjang slice sudah penuh dan kita menambahkan data baru,
	maka function tersebut akan membuat sebuah array baru dan tidak mempengaruhi array sebelumnya. Jika, slice tersebut belum full,
	maka function append tidak membuat array dan dapat mempengaruhi data sebelumnya*/

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 6)
	newSlice[0] = "Andre"
	newSlice[1] = "Johi"

	fmt.Println("New Slice :", newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//Perbedaan Deklarasi Array dan Slice
	iniArray := [...]int{1, 2, 3} //Array perlu mendefenisikan panjangny atau dengan [...]
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

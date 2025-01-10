package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	//var err error = errors.New("Ups Error")
	//fmt.Print(err.Error())

	hasil, err := Pembagi(100, 0)

	if err == nil {
		fmt.Println("Hasil pembagian :", hasil)
	} else {
		fmt.Print("Error :", err.Error())
	}
}

package main

import "fmt"

func main() {
	type KTP string
	type married bool

	var noKTP KTP = "193023"
	var marriedStatus married = false

	fmt.Println(noKTP)
	fmt.Println(marriedStatus)
}

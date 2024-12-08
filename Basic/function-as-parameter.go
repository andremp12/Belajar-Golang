package main

import "fmt"

/*func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Nama : " + nameFiltered)
}*/

// Penggunaan Type Declaration -> untuk argumen function yang kepanjangan
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Nama : " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kasar" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Kasar", spamFilter)
	//Atau

	filter := spamFilter
	sayHelloWithFilter("Andre", filter)
}

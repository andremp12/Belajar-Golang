package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Andre Mahesa Putra",
		"address": "Padang",
	}

	person["age"] = "21"

	fmt.Println("Map Person : ", person)
	fmt.Println("Nama : ", person["name"])
	fmt.Println("Address : ", person["address"])
	fmt.Println("Age : ", person["age"])

	book := make(map[string]string)
	book["title"] = "Romance"
	book["author"] = "Ridho"
	book["publish"] = "Padang"

	delete(book, "publish")

	fmt.Println("\nMap Book : ", book)
}

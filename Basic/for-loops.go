package main

import "fmt"

func main() {
	/*counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-:", counter)
		counter++
	}*/

	//For Loops dengan Statement
	fmt.Println("-For Loops menggunakan statement-")
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	//For Range
	fmt.Println("\n-For Range-")

	slice := []string{"Andre", "Joko", "Zakki"}
	for i, value := range slice {
		fmt.Println("Slice index ", i, " :", value)
	}

	fmt.Println("\n")
	person := make(map[string]string)
	person["name"] = "Andre"
	person["address"] = "Padang"
	person["age"] = "21"

	for key, value := range person {
		fmt.Println(key, " : ", value)
	}

}

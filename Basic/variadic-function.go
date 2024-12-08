package main

import "fmt"

func main() {
	name, total := sumAll("Andre", 10, 10, 5, 7)

	fmt.Println("Name :", name)
	fmt.Println("Total :", total)

	//Slice Parameter
	slice := []int{1, 2, 3, 4, 5, 10}
	name, total = sumAll("Aji", slice...)

	fmt.Println("\nParameter Slice")
	fmt.Println("Name :", name)
	fmt.Println("Total :", total)
}

func sumAll(name string, numbers ...int) (string, int) { // Variadic hanya bisa digunakan dibagian akhir argumen, seperti pada contohn numbers tersebut
	total := 0

	for _, number := range numbers {
		total += number
	}

	return name, total
}

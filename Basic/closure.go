package main

import "fmt"

func main() {
	counter := 0
	name := "Aji"

	increment := func() {
		//name = "Budi"
		name := "Budi"
		fmt.Print("Increment\n")
		counter++
		fmt.Println(name)
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}

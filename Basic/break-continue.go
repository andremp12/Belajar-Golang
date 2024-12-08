package main

import "fmt"

func main() {

	//Penggunaan Break
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}

		fmt.Println("Perulangan ke- ", i)
	}

	//Penggunaan Continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke- ", i)
	}
}

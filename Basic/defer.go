package main

import "fmt"

func logging() {
	fmt.Print("Selesai Memanggil Function Logging\n")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("RunApplication")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}

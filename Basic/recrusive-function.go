package main

import "fmt"

func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	factorialRecrusive := factorial(6)

	fmt.Print(factorialRecrusive)
}

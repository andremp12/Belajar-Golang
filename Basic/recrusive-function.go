package main

func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {

}

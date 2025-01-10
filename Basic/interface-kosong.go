package main

import "fmt"

func TesAngka(i int, any interface{}) interface{} {
	if i == 1 && any == 1 {
		return i
	} else if i == 2 && any == true {
		return true
	} else if i == 3 && any == "3" {
		return any
	} else {
		return "Gagal"
	}
}

func main() {
	var data interface{} = TesAngka(3, 4)

	fmt.Println(data)
}

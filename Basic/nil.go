package main

import "fmt"

func NeWMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	//var person map[string]string = nil
	//fmt.Println(person)

	person := NeWMap("Andre")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}

package main

import (
	"fmt"
)

func random() interface{} {
	return 1
}

func main() {
	result := random()
	//resultString := result.(string)
	//
	//fmt.Println(resultString)

	//Type Assertions Switch

	//switch result.(type) {
	//case string:
	//	fmt.Println("String :", result.(string))
	//case int:
	//	fmt.Println("Integer :", result.(int))
	//default:
	//	fmt.Println("Unknown type")
	//}

	switch value := result.(type) {
	case string:
		fmt.Println("String :", value)
	case int:
		fmt.Println("Integer :", value)
	default:
		fmt.Println("Unknown type")
	}
}

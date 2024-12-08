package main

import "fmt"

func main() {
	/*firstName := "Andre"*/
	name := "Aji"

	/*sayHello()*/
	/*sayHelloTo(firstName, "Putra")*/

	_, lastName, _ := getFullName(name) // Menggunakan _ untuk menghiraukan beberapa return value

	/*fmt.Println(firstName)*/
	/*fmt.Println(nickname)*/
	fmt.Println(lastName)

	firstName, lastName2, age := getCompleteName(21)
	fmt.Println("First Name :", firstName)
	fmt.Println("Last Name :", lastName2)
	fmt.Println("Age : ", age)
}

func sayHello() {
	fmt.Println("Hello")
}

// Function Parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// Function Return Value
func getHello(name string) string {
	return "Hello " + name
}

// Function Returning Multiple Values
func getFullName(name string) (string, string, string) {
	return name, "Baik", "Siap"
}

// Named Return Values
func getCompleteName(ageValue int8) (firstName, lastName string, age int8) {
	firstName = "Eko"
	lastName = "Ramadhan"

	return firstName, lastName, ageValue
}

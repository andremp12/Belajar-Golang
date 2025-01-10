package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
	Age  int
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	var ridho Person
	cat := Animal{
		Name: "cat",
	}

	ridho.Age = 12
	ridho.Name = "Ridho"

	SayHello(ridho)
	SayHello(cat)
}

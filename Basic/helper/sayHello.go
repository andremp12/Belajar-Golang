package helper

import "fmt"

// Jika penamaan diawali huruf besar, bisa di akses di package lain
var Application = "golang"

// Jika penamaa diawali huruf kecil, tidak bisa di akses/digunakan di package lain
var version = "1.0.0"

// Bisa diakses di package lain
func SayHello() {
	fmt.Printf("Hello")
}

// Hanya bisa di dalam satu package saat ini
func sayHello(name string) {
	fmt.Printf("Hello", name)
}

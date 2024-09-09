package main

import "fmt"

func main() {
	name := "Maulana"

	if name == "imam" {
		fmt.Println("Hello imam")
	} else if name == "Akbar" {
		fmt.Println("Hello Akbar")
	} else if name == "Archeus" {
		fmt.Println("Hello Archeus")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}

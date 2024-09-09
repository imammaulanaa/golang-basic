package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Imam")
	fmt.Println(result)

	fmt.Println(getHello("Maulana"))
	fmt.Println(getHello("Akbar"))
}

package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Imam"
	middleName = "Maulana"
	lastName = "Akbar"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

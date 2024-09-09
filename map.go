package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Imam"
	//person["address"] = "Bekasi"

	person := map[string]string{
		"name":    "Imam",
		"address": "Bekasi",
	}

	fmt.Println(person["name"])
	fmt.Println(person["Bekasi"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Golang Basic"
	book["author"] = "Imam"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}

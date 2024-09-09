package main

import (
	"fmt"
	"golang-basic/helper"
)

func main() {
	result := helper.SayHello("Imam")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Imam")) // tidak bisa diakses
}

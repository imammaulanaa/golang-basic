package main

import "fmt"

func main() {

	type NoKTP string

	var ktpimam NoKTP = "1111111111"

	var contoh string = "222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpimam)
	fmt.Println(contohKtp)

}

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var imam Customer
	fmt.Println(imam)

	imam.Name = "Imam Maulana"
	imam.Address = "Indonesia"
	imam.Age = 30
	fmt.Println(imam)
	fmt.Println(imam.Name)
	fmt.Println(imam.Address)
	fmt.Println(imam.Age)

	akbar := Customer{
		Name:    "akbar",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(akbar)

	agus := Customer{"agus", "Indonesia", 30}
	fmt.Println(agus)

	akbar.sayHello("akbar")
	imam.sayHello("imam")
	agus.sayHello("Agus")
}

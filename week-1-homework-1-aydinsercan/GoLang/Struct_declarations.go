package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name    string
	surname string
	age     int
}

func (p Person) greet() string {
	str := strconv.Itoa(p.age)
	return "Selam " + p.name + " :) " + p.surname + " " + str
}

func main() {
	var greeter Person = Person{"Nihal", "Aydın", 27}
	greeting := greeter.greet()
	fmt.Printf("%s\n", greeting)
}

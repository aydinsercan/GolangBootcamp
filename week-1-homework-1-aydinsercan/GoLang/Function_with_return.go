package main

import (
	"fmt"
)

func main() {
	name := "Betül"
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}

package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	greeting := createGreet(name)
	greet(name)

	fmt.Printf("%s\n", greeting)
}

func greet(name string) {
	fmt.Printf("Hello %s :)\n", name)
}

func createGreet(name string) string {
	return "Selam " + name + " :)"
}

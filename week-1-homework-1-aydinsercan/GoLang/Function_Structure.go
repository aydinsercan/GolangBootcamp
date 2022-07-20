package main

import (
	"fmt"
)

func main() {
	greet("Sercan")
}

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}

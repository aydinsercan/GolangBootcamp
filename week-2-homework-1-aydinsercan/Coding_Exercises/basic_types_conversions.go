//An exercise for conversions among basic types

package main

import "fmt"

func main() {

	func1()
	func2()

}

//Basic type conversions
func func1() {
	x := 100
	y := 2.5

	fmt.Printf("x value     : %T\n", x)
	fmt.Printf("y value		: %T\n", float64(y))
	fmt.Printf("mul value	: %T\n", float64(x)*y)

	mul := int(float64(x) * y)

	fmt.Printf("Mul Value: %d", mul)
}

func func2() {
	var men uint8
	men = 5
	var women int16
	women = 6

	var people int

	people = int(men) + int(women)

	fmt.Printf("People: %d", people)

}

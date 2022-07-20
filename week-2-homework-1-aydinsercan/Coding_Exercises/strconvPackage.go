package main

//STRCONV PACKAGE BUILT IN METHODS
import (
	"fmt"
	"strconv"
)

func main() {

	FuncItoa()
	FuncAtoi()
	FuncParseFloat()
}

func FuncAtoi() {
	// Using Atoi() function
	x1 := "245"
	fmt.Println("Before:")
	fmt.Printf("Type: %T ", x1)
	fmt.Printf("\nValue: %v", x1)
	y1, e1 := strconv.Atoi(x1)
	if e1 == nil {
		fmt.Println("\nAfter:")
		fmt.Printf("Type: %T ", y1)
		fmt.Printf("\nValue: %v", y1)
	}

	x2 := "1"
	fmt.Println("\n\nBefore:")
	fmt.Printf("Type: %T ", x2)
	fmt.Printf("\nValue: %v", x2)
	y2, e2 := strconv.Atoi(x2)
	if e2 == nil {
		fmt.Println("\nAfter:")
		fmt.Printf("Type: %T ", y2)
		fmt.Printf("\nValue: %v", y2)
	}
}

func FuncItoa() {
	num := 200
	numStr := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", numStr, numStr)
}

func FuncParseFloat() {
	strF := "250.56"
	str, err := strconv.ParseFloat(strF, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type:%T value:%#v\n", str, str)
}

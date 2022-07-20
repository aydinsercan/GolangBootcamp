package main

import "fmt"

var bbb = "test2"

//SCOPE
func testGlobal() {
	fmt.Println(bbb)
}

func main() {
	var aaa = "test"
	fmt.Println(aaa)
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	testLocal()
	testGlobal()
}

func testLocal() {
	//fmt.Println(aaa) //Undefined here
}

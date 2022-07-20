package main

import "fmt"

type baseGroup int

const (
	fooGroup baseGroup = iota + 1
	barGroup
)

//IOTA
var groups = [...]string{fooGroup: "foo", barGroup: "bar"}

var xGroups = map[baseGroup]string{fooGroup: "foo", barGroup: "bar"}

func main() {
	fmt.Println("groups")
	for k, v := range groups {
		fmt.Println(k, v)
	}

	fmt.Println("xGroups")
	for k, v := range xGroups {
		fmt.Println(k, v)
	}
}

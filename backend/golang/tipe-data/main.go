package main

import (
	"fmt"

)

func main() {
	// var name string = "hilmi"
	// name = "hilmi 2"

	// var name2 = "waky"
	// name3 := "vestia"

	// fmt.Println(name, name2, name3)

	// =========================================================

	var x, y int64

	x = 100
	y = 5

	result := x % y
	var isGenap bool
	if result == 0 {
		isGenap = true
	} else {
		isGenap = false
	}

	if isGenap {
		fmt.Println("bilangan genap")

	} else {
		fmt.Println("bilangan ganjil")
	}

	fmt.Println(result)
}

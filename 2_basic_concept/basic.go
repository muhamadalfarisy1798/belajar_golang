package main

import (
	"fmt"
)

func main() {
	var i, j int = 2, 6
	fmt.Println(i, j)
	// taking input
	var input int
	fmt.Scanln("masukkan angka : ", &input)
	// & to return the address of a variable
	fmt.Println(input)

	// conditional need a variable
	if x := 42; x > 18 {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Restricted")
	}

	num := 3
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}

	// loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	res := 0
	for i := 0; i < 10; i += 3 {
		res += i
	}
	fmt.Println(res)

	// Quiz 1

}

func number(num int) {
	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	}
}

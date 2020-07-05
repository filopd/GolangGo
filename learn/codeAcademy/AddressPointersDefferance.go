package main

import "fmt"

func double(x int) {
	x += x
}

func double_1(x *int) {
	*x += *x
}

func main() {
	// Problem of scope of vars in methods/blocks.
	varX := 100
	double(varX)
	fmt.Println("varX is", varX) // varX is 100
	// Address, Pointer and Defferencing
	varA := "Lol1"
	var varAPointer *string
	varAPointer = &varA //varAPointer := &varA
	fmt.Println("Variable is varA with value", varA)
	fmt.Println("varA Address is (&varA)", &varA)
	fmt.Println("varAPointer Pointer is", varAPointer)
	*varAPointer = "Pop2"
	fmt.Println("New varA by changing *varAPointer", varA)
	// Solution to Problem
	double_1(&varX)
	fmt.Println("varX is", varX) // varX is 200
}

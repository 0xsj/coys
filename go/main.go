package main

import (
	"fmt"
	"math"
)

// simple hello world
func sayHello() {
	// fmt.Println("hello world")
}

func printValues() {
	// fmt.Println("go" + "lang")
	// fmt.Println("1 + 1 = ", 1+1)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)
}

func variables() {

	// declaring a single variable
	// var a = "intial"
	// fmt.Println(a)

	// // multi declaration
	// var b, c int = 1, 2
	// fmt.Println(b + c)

	// // type inference, as we don't explicitly say bool
	// var d = true
	// fmt.Println(d)

	// // variables with no value for an int is 0
	// var e int
	// fmt.Println(e)

	// // := is shorthand for declaring and initializing a variable
	// f := "apple"
	// fmt.Println(f)
}

func constants() {
	// constant declares a variable
	const s string = "constant"
	fmt.Println(s)

	// you can have constants inside and outside of function scope
	const n = 500000000

	//
	const d = 3e20 / n
	fmt.Println(d)

	// constant expressions can be converted to a type
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func main() {
	fmt.Println("main running")
	sayHello()
	printValues()
	variables()

	constants()
}

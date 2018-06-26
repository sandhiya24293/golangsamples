package main

import (
"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	//Closures
	//Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.
	//An example will make things more clear.
	x := 5
	func() {
		fmt.Println("x =", x)
	}()
	//In the program above, the anonymous function accesses the variable a which is present outside it's body in line no. 22.
	// Hence this anonymous function is a closure.
	// Every closure is bound to its own surrounding variabl.
	// Let's understand what this means by using a simple example.

	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
	//In the program above, the function appendStr returns a closure.
	// This closure is bound to the variable t. Let's understand what this means.
	//The variables a and b declared in line nos. 29, 30 are closures and they are bound to their own value of t.
	//We first call a with the parameter World.
	// Now the value of a's version of t becomes Hello World.
	//In line no. 32 we call b with the parameter Everyone.
	// Since b is bound to it's own variable t, b's version of t has a initial value of Hello again. Hence after this function call, the value of b's version of t becomes Hello Everyone. The rest of the program is self explanatory.

	//This program will print,
	/*
	Hello World
	Hello Everyone
	Hello World Gopher
	Hello Everyone !
	*/
}

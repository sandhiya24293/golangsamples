package main

import "fmt"


func change(val *int) {
	*val = 55
}

func main() {
	//What is a pointer?
	//A pointer is a variable which stores the memory address of another variable.
	//Go does not support pointer arithmetic which is present in other languages like C.


	//Declaring pointers
	//*T is the type of the pointer variable which points to a value of type T.
	b := 255
	var a *int = &b		//The & operator is used to get the address of a variable.
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//The zero value of a pointer is nil.
	x := 25
	var y *int
	if b == nil {
		fmt.Println("b is", y)
		y = &x
		fmt.Println("b after initialization is", b)
	}

	//Dereferencing a pointer
	//Dereferencing a pointer means accessing the value of the variable which the pointer points to.
	//*a is the syntax to deference a.

	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)

	//change the value in b using the pointer.
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)

	//Passing pointer to a function
	a1 := 58
	fmt.Println("value of a1 before function call is",a1)
	b1 := &a1
	change(b1)
	fmt.Println("value of a after function call is", a1)

	//Go does not support pointer arithmetic
	/*b := [...]int{109, 110, 111}
	p := &b
	p++*/

	//Excercise
	//1) Change the value by Passing pointer to a function
	//2) Do not pass a pointer to an array as a argument to a function. Use slice instead why?


}

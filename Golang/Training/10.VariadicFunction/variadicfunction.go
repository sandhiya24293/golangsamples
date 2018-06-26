package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}


func main(){
	//What is a variadic function?
	//A variadic function is a function that can accept variable number of arguments.

	//Syntax
	//If the last parameter of a function is denoted by ...T, then the function can accept any number of arguments of type T for the last parameter.

	//Please note that only the last parameter of a function is allowed to be variadic.

	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	/*The way variadic functions work is by converting the variable number of arguments passed,
	to a new slice of the type of the variadic parameter.
	For instance in line no. 22 of the program above,
	the variable number of arguments to the find function are 89, 90, 95.
	The find function expects a variadic int argument.
	Hence these three arguments will be converted by the compiler to a slice of type int []int{89, 90, 95}
	and then it will be passed to the find function.
	*/

	//Passing a slice to a variadic function
	//Try tho run the below commented code.
	/*
	nums := []int{89, 90, 95}
	find(89, nums)
	*/
	//This will not work.
	// The above program will fail with compilation error main.go:46: cannot use nums (type []int) as type int in argument to find
	//Why does this doesn't work?
	//As we already discussed, these variadic arguments will be converted to a slice of type int since find expects variadic int arguments.
	// In this case nums is already a int slice and a new []int slice is tried to be created using nums i.e the compiler tries to do
	//find(89, []int{nums})
	//which will fail since nums is a []int and not an int.

	//There is a syntactic sugar which can be used to pass a slice to a variadic function.
	// You have to suffix the slice with ... If that is done, the slice is directly passed to the function without a new slice being created.
	nums := []int{89, 90, 95}
	find(89, nums...)









	//EXCERCISE
	//1) Passing a slice to a variadic function?

}

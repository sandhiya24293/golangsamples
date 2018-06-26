//Gotcha
//Just be sure you know what you are doing when you are modifying a slice inside a variadic function.

//Lets look at a simple example.

package main

import (
	"fmt"
)

func changeA(s ...string) {
	s[0] = "Go"
}

func changeB(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {

	welcomeA := []string{"hello", "world"}
	changeA(welcomeA...)
	fmt.Println(welcomeA)

	//What do you think will be the output of the above program?
	// If you think it will be [Go world] Congrats! you have understood variadic functions and slices.
	// If you got it wrong, no big deal, let me explain how we get this output.
	//
	//In line no. 18 of the program above, we are using the syntactic sugar ... and passing the slice as a variadic argument to the changeA function.
	//As we have already discussed, if ... is used, the welcome slice itself will be passed as an argument without a new slice being created.
	// Hence welcome will be passed to the changeA function as argument.
	//Inside the changeA function, the first element of the slice is changed to Go. Hence this program outputs
	//[Go world]

	//Here is one more program to understand variadic functions.
	welcomeB := []string{"hello", "world"}
	changeB(welcomeB...)
	fmt.Println(welcomeB)
	//Try to figure out how the above program works :).
}




package main

import "fmt"

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func fullNameDefer(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func main() {

	//What is panic?
	//The idiomatic way to handle abnormal conditions in a program in Go is using errors.
	// Errors are sufficient for most of the abnormal conditions arising in the program.
	//But there are some situations where the program cannot simply continue executing after an abnormal situation.
	// In this case we use panic to terminate the program.
	// When a function encounters a panic, it's execution is stopped, any deferred functions are executed and then the control returns to it's caller.
	// This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates.

	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	//When should panic be used?
	//One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.
	//There are two valid use cases for panic.
	//1) An unrecoverable error where the program cannot simply continue its execution.
		//One example would be a web server which fails to bind to the required port. In this case it's reasonable to panic as there is nothing else to do if the port binding itself fails.
	//2) A programmer error.
		//Let's say we have a method which accepts a pointer as a parameter and someone calls this method using nil as argument. In this case we can panic as its a programmer error to call a method with nil argument which was expecting a valid pointer.

	//Defer while panicking
	// When a function encounters a panic, it's execution is stopped, any deferred functions are executed and then the control returns to it's caller.
	// This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates.

	defer fmt.Println("deferred call in main")
	firstName1 := "Elon"
	fullNameDefer(&firstName1, nil)
	fmt.Println("returned normally from main")



}

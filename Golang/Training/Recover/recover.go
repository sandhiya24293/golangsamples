package main

import "fmt"

func recoverName() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
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
	//recover is a builtin function which is used to regain control of a panicking goroutine.
	//The signature of recover function is provided below,
	/*
		func recover() interface{}
	*/
	//Recover is useful only when called inside deferred functions. Executing a call to recover inside a deferred function stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function, it will not stop a panicking sequence.

	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")


	//EXCERCISE
	//1) How to get stack trace after recover?


	//Keypoints for Interview
	//=================//
	//What is panic?
	//When should panic be used?
	//Panic example
	//Defer while panicking
	//Recover
	//Panic, Recover and Goroutines
	//Runtime panics
	//Getting stack trace after recover
}
